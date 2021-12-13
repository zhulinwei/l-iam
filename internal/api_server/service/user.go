package service

import (
	"context"
	"l-iam/internal/api_server/dao"
	v1 "l-iam/internal/pkg/model/api_server/v1"
	"strings"
)

type UserService interface {
	Delete(ctx context.Context, userId uint) error
	Create(ctx context.Context, user *v1.User) error
	Get(ctx context.Context, userId uint) (*v1.User, error)
	Update(ctx context.Context, userId uint, update *v1.User) error
	List(ctx context.Context, opts *v1.UserQueryOptions) (*v1.UserList, error)
}

var _ UserService = (*userService)(nil)

type userService struct {
	dao dao.Factory
}

func newUserService(dao dao.Factory) UserService {
	return &userService{dao: dao}
}

func (s *userService) Get(ctx context.Context, userId uint) (*v1.User, error) {
	user, err := s.dao.Users().Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	user.Password = strings.Replace(user.Password, user.Password[1:], "***", 4)
	return user, nil
}

func (s *userService) Create(ctx context.Context, user *v1.User) error {
	return s.dao.Users().Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, userId uint, update *v1.User) error {
	user, err := s.Get(ctx, userId)
	if err != nil {
		return err
	}

	// 按需更新
	if update.Email != "" {
		user.Email = update.Email
	}
	if update.Phone != "" {
		user.Phone = update.Phone
	}
	if update.Username != "" {
		user.Username = update.Username
	}
	if update.Password != "" {
		user.Password = update.Password
	}

	return s.dao.Users().Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, userId uint) error {
	return s.dao.Users().Delete(ctx, userId)
}

func (s *userService) List(ctx context.Context, opts *v1.UserQueryOptions) (*v1.UserList, error) {
	userList, err := s.dao.Users().List(ctx, opts)
	if err != nil {
		return nil, err
	}
	for _, user := range userList.List {
		user.Password = strings.Replace(user.Password, user.Password[1:], "***", 4)
	}
	return userList, nil
}
