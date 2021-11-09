package dao

import (
	"context"
	"l-iam/internal/api_server/model"

	"gorm.io/gorm"
)

// UserDao defines the user interface.
type UserDao interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, username string) error
	BatchDelete(ctx context.Context, usernames []string) error
	Get(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context, options interface{}) (*model.UserList, error)
}

type userDao struct {
	db *gorm.DB
}

var _ UserDao = (*userDao)(nil)

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}

func (d *userDao) Create(ctx context.Context, user *model.User) error {
	return nil
}

func (d *userDao) Update(ctx context.Context, user *model.User) error {
	return nil
}

func (d *userDao) Delete(ctx context.Context, username string) error {
	return nil
}

func (d *userDao) BatchDelete(ctx context.Context, usernames []string) error {
	return nil
}

func (d *userDao) Get(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}

func (d *userDao) List(ctx context.Context, options interface{}) (*model.UserList, error) {
	return nil, nil
}
