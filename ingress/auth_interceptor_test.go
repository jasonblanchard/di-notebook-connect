package ingress

import (
	"context"
	"fmt"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
)

func TestAuthInterceptor(t *testing.T) {
	idToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMDMxNTY2NTIxNjA3MjU5NTUzOTkifQ.Bi89hZS5TPlPYNvgeE6MifFHP7DpeOiRJeSl79nh8c4"
	interceptor := NewAuthInterceptor()

	unaryFunc := func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		return &connect.Response[any]{}, nil
	}

	wrapped := interceptor(unaryFunc)

	ctx := context.TODO()

	req := &connect.Request[any]{}

	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", idToken))

	_, err := wrapped(ctx, req)

	assert.Nil(t, err)

	assert.Equal(t, "2b5545ef-3557-4f52-994d-daf89e04c390", req.Header().Get("x-principal-id"))
}

func TestAuthInterceptorNoHeader(t *testing.T) {
	interceptor := NewAuthInterceptor()

	unaryFunc := func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		return &connect.Response[any]{}, nil
	}

	wrapped := interceptor(unaryFunc)

	ctx := context.TODO()

	req := &connect.Request[any]{}

	_, err := wrapped(ctx, req)

	assert.Equal(t, "unauthenticated: no Authorization header", err.Error())
}

func TestAuthInterceptorNonsenseToken(t *testing.T) {
	idToken := "nonsense"
	interceptor := NewAuthInterceptor()

	unaryFunc := func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		return &connect.Response[any]{}, nil
	}

	wrapped := interceptor(unaryFunc)

	ctx := context.TODO()

	req := &connect.Request[any]{}

	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", idToken))

	_, err := wrapped(ctx, req)

	assert.Equal(t, "unauthenticated: no sub in bearer token", err.Error())

}
