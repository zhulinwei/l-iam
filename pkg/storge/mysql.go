package storge

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DefaultCharset            = "utf8mb4"
	DefaultParseTime          = "True"
	DefaultLocation           = "Local"
	DefaultOpenConnections    = 100
	DefaultIdleConnections    = 100
	DefaultConnectionLiftTime = time.Second * 10
)

type MySQLOptions struct {
	Address               string        `json:"address"`
	Username              string        `json:"username"`
	Password              string        `json:"-"`
	Database              string        `json:"database"`
	Charset               string        `json:"charset"`
	Location              string        `json:"location"`
	ParseTime             string        `json:"parse_time"`
	LogLevel              int           `json:"log_level"`
	MaxIdleConnections    int           `json:"max_idle_connections"`
	MaxOpenConnections    int           `json:"max_open_connections"`
	MaxConnectionLifeTime time.Duration `json:"max_connection_life_time"`
}

func New(options MySQLOptions) (*gorm.DB, error) {
	var dsn string

	if options.Charset == "" {
		options.Charset = DefaultCharset
	}
	if options.ParseTime == "" {
		options.ParseTime = DefaultParseTime
	}
	if options.Location == "" {
		options.Location = DefaultLocation
	}
	if options.MaxIdleConnections == 0 {
		options.MaxIdleConnections = DefaultIdleConnections
	}
	if options.MaxOpenConnections == 0 {
		options.MaxOpenConnections = DefaultOpenConnections
	}
	if options.MaxConnectionLifeTime < time.Second {
		options.MaxConnectionLifeTime = DefaultConnectionLiftTime
	}

	if options.Username == "" && options.Password == "" {
		dsn = fmt.Sprintf("tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s",
			options.Address, options.Database, options.Charset, options.ParseTime, options.Location)
	} else if options.Username != "" && options.Password != "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s",
			options.Username, options.Password, options.Address, options.Database, options.Charset, options.ParseTime, options.Location)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(options.MaxIdleConnections)
	// 最大打开的连接数
	sqlDB.SetMaxOpenConns(options.MaxOpenConnections)
	// 空闲连接最大的存活时间
	sqlDB.SetConnMaxLifetime(options.MaxConnectionLifeTime)

	return db, nil
}
