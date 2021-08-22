package api_user_v2

import (
	// Internal
	"context"
	"errors"

	// Internal
	user_v2 "github.com/iakrevetkho/archaeopteryx/proto/user/v2"
)

var (
	WRONG_USER_ID_ERROR = errors.New("wrong user id")
)

func (s *UserServiceServer) GetUser(ctx context.Context, request *user_v2.GetUserRequest) (*user_v2.GetUserResponse, error) {
	if request.GetId() == 0 {
		return nil, WRONG_USER_ID_ERROR
	}
	return &user_v2.GetUserResponse{FirstName: "Bobby", LastName: "Twist", Password: "qwerty"}, nil
}
