package dao

import (
	"l-iam/internal/api_server/config/options"
	v1 "l-iam/internal/pkg/model/api_server/v1"
	"l-iam/pkg/log"
	"l-iam/pkg/storge"
	"sync"

	"gorm.io/gorm"
)

// ------------- 定义dao接口 ----------

var client Factory

type Factory interface {
	Users() UserDao
	Policies() PolicyDao
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}

// -------------- mysql实现-----------

// apiServerFactory 使用MySQL对IFactory的实现
type apiServerFactory struct {
	db *gorm.DB
}

var once sync.Once

func NewApiServerFactory(option *options.MySQLOptions) (Factory, error) {
	var err error
	var db *gorm.DB

	once.Do(func() {
		db, err = storge.New(storge.MySQLOptions{
			Address:               option.Host,
			Username:              option.Username,
			Password:              option.Password,
			Database:              option.Database,
			MaxOpenConnections:    option.MaxOpenConnections,
			MaxIdleConnections:    option.MaxIdleConnections,
			MaxConnectionLifeTime: option.MaxConnectionLifeTime,
		})

	})

	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&v1.Policy{})
	if err != nil {
		log.Info(err.Error())
	}
	return &apiServerFactory{db: db}, nil
}

func (a *apiServerFactory) Users() UserDao {
	return NewUserDao(a.db)
}

func (a *apiServerFactory) Policies() PolicyDao {
	return NewPolicyDao(a.db)
}
