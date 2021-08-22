package api_user_v1

import (
	// External
	"context"
	"errors"

	// Internal
	user_v1 "github.com/iakrevetkho/archaeopteryx/proto/user/v1"
)

var (
	WRONG_USER_ID_ERROR = errors.New("wrong user id")
)

func (s *UserServiceServer) GetUser(ctx context.Context, request *user_v1.GetUserRequest) (*user_v1.GetUserResponse, error) {
	if request.GetId() == 0 {
		return nil, WRONG_USER_ID_ERROR
	}
	return &user_v1.GetUserResponse{Name: "Bobby", Password: "qwerty"}, nil
}
