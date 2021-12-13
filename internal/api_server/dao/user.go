package dao

import (
	"context"
	v1 "l-iam/internal/pkg/model/api_server/v1"
	"l-iam/pkg/code"
	"l-iam/pkg/errors"

	"gorm.io/gorm"
)

// UserDao defines the user interface.
type UserDao interface {
	Create(ctx context.Context, user *v1.User) error
	Get(ctx context.Context, userId uint) (*v1.User, error)
	Update(ctx context.Context, user *v1.User) error
	Delete(ctx context.Context, userId uint) error
	List(ctx context.Context, opts *v1.UserQueryOptions) (*v1.UserList, error)
}

type userDao struct {
	db *gorm.DB
}

var _ UserDao = (*userDao)(nil)

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}

func (d *userDao) Create(ctx context.Context, user *v1.User) error {
	err := d.db.WithContext(ctx).Create(&user).Error
	if errors.IsMysqlError1062(err) {
		return errors.WithCode(code.ErrUserAlreadyExist, "User %s alert exist", user.Username)
	}
	return err
}

func (d *userDao) Update(ctx context.Context, user *v1.User) error {
	return d.db.WithContext(ctx).Updates(user).Error
}

func (d *userDao) Delete(ctx context.Context, userId uint) error {
	return d.db.WithContext(ctx).Delete(&v1.User{}, userId).Error
}

func (d *userDao) Get(ctx context.Context, userId uint) (*v1.User, error) {
	user := &v1.User{}

	err := d.db.WithContext(ctx).Where("id = ?", userId).First(user).Error
	switch err {
	case nil:
		return user, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, errors.WrapError(code.ErrUserNotFount, err)
	}
}

func (d *userDao) List(ctx context.Context, opts *v1.UserQueryOptions) (*v1.UserList, error) {
	db := d.buildQueryDB(ctx, opts)
	res := &v1.UserList{}

	err := db.Find(&res.List).Offset(-1).Limit(-1).Count(&res.Total).Error
	switch err {
	case nil:
		return res, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, errors.WrapError(code.ErrUserNotFount, err)
	}
}

func (d *userDao) buildQueryDB(ctx context.Context, opts *v1.UserQueryOptions) *gorm.DB {
	db := d.db.WithContext(ctx).Model(&v1.User{})
	if opts.Page > 0 && opts.Size > 0 {
		db = db.Offset((opts.Page - 1) * opts.Size).Limit(opts.Size)
	}
	if opts.Name != "" {
		db = db.Where("username like ?", "%"+opts.Name+"%")
	}
	return db.Order("id desc")
}
