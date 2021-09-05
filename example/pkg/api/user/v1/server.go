package api_user_v1

import (
	// External
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	// Internal
	user_v1 "github.com/iakrevetkho/archaeopteryx/example/proto/gen/user/v1"
	api_data "github.com/iakrevetkho/archaeopteryx/pkg/api/data"
	"github.com/iakrevetkho/archaeopteryx/pkg/helpers"
)

type UserServiceServer struct {
	log *logrus.Entry
	// Required to have revese compatability
	user_v1.UnimplementedUserServiceServer
}

func RegisterServiceServer(s grpc.ServiceRegistrar, controllers *api_data.Controllers) error {
	server := new(UserServiceServer)
	server.log = helpers.CreateComponentLogger("grpc-user-v1")

	// Attach the User service to the server
	user_v1.RegisterUserServiceServer(s, server)
	server.log.Debug("Service registered")

	return nil
}

func RegisterProxyServiceServer(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	// Attach handler to global handler
	if err := user_v1.RegisterUserServiceHandler(context.Background(), mux, conn); err != nil {
		return err
	}

	return nil
}
