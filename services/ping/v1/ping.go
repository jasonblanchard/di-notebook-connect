package pingv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	pingv1 "github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/ping/v1"
)

type Service struct{}

func (s *Service) Ping(ctx context.Context, req *connect.Request[pingv1.PingRequest]) (*connect.Response[pingv1.PingResponse], error) {
	res := connect.NewResponse(&pingv1.PingResponse{
		Sound: "pong",
	})
	return res, nil
}
