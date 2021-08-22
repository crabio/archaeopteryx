package api_user_v1

import (
	// External
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	// Internal
	user_v1 "github.com/iakrevetkho/archaeopteryx/proto/user/v1"
)

type UserServiceServer struct {
	// Required to have revese compatability
	user_v1.UnimplementedUserServiceServer
}

func RegisterServiceServer(s grpc.ServiceRegistrar) error {
	// Attach the User service to the server
	user_v1.RegisterUserServiceServer(s, &UserServiceServer{})

	return nil
}

func RegisterProxyServiceServer(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	// Attach handler to global handler
	if err := user_v1.RegisterUserServiceHandler(context.Background(), mux, conn); err != nil {
		return err
	}
	return nil
}
