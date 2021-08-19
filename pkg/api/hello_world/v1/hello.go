package api_hello_world_v1

import (
	"context"

	hello_world_v1 "github.com/iakrevetkho/archaeopteryx/proto/hello_world/v1"
)

func (s *HelloServiceServer) SayHello(ctx context.Context, in *hello_world_v1.SayHelloRequest) (*hello_world_v1.SayHelloResponse, error) {
	return &hello_world_v1.SayHelloResponse{Message: in.Name + " world"}, nil
}
