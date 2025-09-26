package handler

import (
	"context"
	"errors"

	userpb "freedom/gen/proto/user"

	"github.com/google/uuid"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名和密码不能为空")
	}

	// 模拟生成用户 ID
	userID := uuid.New().String()

	return &userpb.RegisterResponse{
		UserId:   userID,
		Username: req.Username,
		Success:  true,
	}, nil
}
