package api_user_v2

import (
	// External
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	// Internal
	user_v2 "github.com/iakrevetkho/archaeopteryx/proto/user/v2"
)

type UserServiceServer struct {
	// Required to have revese compatability
	user_v2.UnimplementedUserServiceServer
}

func RegisterServiceServer(s grpc.ServiceRegistrar) error {
	// Attach the User service to the server
	user_v2.RegisterUserServiceServer(s, &UserServiceServer{})

	return nil
}

func RegisterProxyServiceServer(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	// Attach handler to global handler
	if err := user_v2.RegisterUserServiceHandler(context.Background(), mux, conn); err != nil {
		return err
	}
	return nil
}
