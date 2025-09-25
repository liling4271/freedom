package service

import (
	"context"
	"errors"
	userpb "freedom/gen/proto/user"
	userdb "freedom/internal/user/model"
	userRepo "freedom/internal/user/repository"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	repo userRepo.UserRepository
}

// 新用户注册时，last_login 应该是 NULL（从未登录过）
func (s *UserService) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
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
	res := &userpb.RegisterResponse{}
	return res, nil
}

func NewUserService(repo userRepo.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
