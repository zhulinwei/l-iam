package dao

import (
	"l-iam/internal/api_server/config/options"
	"l-iam/pkg/storge"
	"sync"

	"gorm.io/gorm"
)

// ------------- 定义dao接口 ----------

var client IFactory

type IFactory interface {
	Users() IUser
	Policies() IPolicy
}

func Client() IFactory {
	return client
}

func SetClient(factory IFactory) {
	client = factory
}

// -------------- mysql实现-----------

// apiServerFactory 使用MySQL对IFactory的实现
type apiServerFactory struct {
	db *gorm.DB
}

var once sync.Once

func NewApiServerFactory(option *options.MySQLOptions) (IFactory, error) {
	var err error
	var db *gorm.DB

	once.Do(func() {
		db, err = storge.New(storge.MySQLOptions{
			Address:  option.Host,
			Username: option.Username,
			Password: option.Password,
			Database: option.Database,
		})
	})

	if err != nil {
		return nil, err
	}

	return &apiServerFactory{db: db}, nil
}

func (a *apiServerFactory) Users() IUser {
	return NewUserDao(a.db)
}

func (a *apiServerFactory) Policies() IPolicy {
	return NewPolicyDao(a.db)
}
