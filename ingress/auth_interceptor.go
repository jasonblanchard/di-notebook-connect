package ingress

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/dgrijalva/jwt-go"
)

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			authHeader := req.Header().Get("Authorization")

			if authHeader == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no Authorization header"))
			}

			sub, err := bearerHeaderToSub(authHeader)

			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no sub in bearer token"))
			}

			principalId, err := getPrincipalIdBySub("google", sub)

			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no principal for sub"))
			}

			req.Header().Set("x-principal-id", principalId)

			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

func bearerHeaderToSub(header string) (string, error) {
	tokenString := strings.Replace(header, "Bearer ", "", 1)

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})

	if token == nil {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("token does not contain any claims")
	}

	return fmt.Sprintf("%s", claims["sub"]), nil
}

func getPrincipalIdBySub(issuer, sub string) (string, error) {
	googleSubs := map[string]string{
		"103156652160725955399": "2b5545ef-3557-4f52-994d-daf89e04c390",
	}

	id, ok := googleSubs[sub]

	if !ok {
		return "", errors.New("no sub")
	}

	return id, nil
}
