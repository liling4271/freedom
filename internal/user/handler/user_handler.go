package handler

import (
	"context"
	"errors"
	"freedom/internal/user/service"

	"google.golang.org/grpc"

	userpb "freedom/gen/proto/user"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

// 注册 gRPC 服务
func RegisterUserService(grpcServer *grpc.Server, handler *UserHandler) {
	userpb.RegisterUserServiceServer(grpcServer, handler)
}

func (h *UserHandler) Register(
	ctx context.Context, // gRPC 要求的 context.Context
	req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) { // gRPC 生成的响应结构体和错误返回
	// 1. 简单参数校验（避免空请求）
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	// 2. 调用服务层处理业务逻辑（将 gRPC 请求传递给服务层）
	resp, err := h.svc.Register(ctx, req)
	if err != nil {
		return nil, errors.Join(errors.New("gRPC 处理注册失败"), err)
	}

	// 3. 返回 gRPC 响应
	return resp, nil
}
