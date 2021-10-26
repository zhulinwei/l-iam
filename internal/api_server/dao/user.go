package dao

import (
	"context"
	"l-iam/internal/api_server/model"

	"gorm.io/gorm"
)

// IUser defines the user interface.
type IUser interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, username string) error
	BatchDelete(ctx context.Context, usernames []string) error
	Get(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context, options interface{}) (*model.UserList, error)
}

var _ IUser = (*UserDao)(nil)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) IUser {
	return &UserDao{db: db}
}

func (d *UserDao) Create(ctx context.Context, user *model.User) error {
	return nil
}

func (d *UserDao) Update(ctx context.Context, user *model.User) error {
	return nil
}

func (d *UserDao) Delete(ctx context.Context, username string) error {
	return nil
}

func (d *UserDao) BatchDelete(ctx context.Context, usernames []string) error {
	return nil
}

func (d *UserDao) Get(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}

func (d *UserDao) List(ctx context.Context, options interface{}) (*model.UserList, error) {
	return nil, nil
}
