package service

import (
	"context"
	"errors"
	userPb "freedom/gen/proto/user"
	userdb "freedom/internal/user/model"
	userRepo "freedom/internal/user/repository"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	// 嵌入未实现的方法
	userPb.UnimplementedUserServiceServer
	repo userRepo.UserRepository
}

// 确保实现了 gRPC 接口
var _ userPb.UserServiceServer = (*UserService)(nil)

func NewUserService(repo userRepo.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// 新用户注册时，last_login 应该是 NULL（从未登录过）
func (s *UserService) Register(ctx context.Context, req *userPb.RegisterRequest) (*userPb.RegisterResponse, error) {
	if req.Username == "" {
		return nil, errors.New("账号不能为空")
	}
	if req.Password == "" {
		return nil, errors.New("密码不能为空")
	}
	user := &userdb.User{
		ID:         uuid.New().String(),
		Username:   req.Username,
		UserStatus: 1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		LastLogin:  nil, // 明确表示从未登录
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	res := &userPb.RegisterResponse{}
	return res, nil
}
