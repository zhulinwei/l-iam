package storge

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DefaultCharset   = "utf8mb4"
	DefaultParseTime = "True"
	DefaultLocation  = "Local"
)

type MySQLOptions struct {
	Address   string `json:"address"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	Charset   string `json:"charset"`
	Location  string `json:"location"`
	ParseTime string `json:"parse_time"`
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

	if options.Username == "" && options.Password == "" {
		dsn = fmt.Sprintf("tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s",
			options.Address, options.Database, options.Charset, options.ParseTime, options.Location)
	} else if options.Username != "" && options.Password != "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s",
			options.Username, options.Password, options.Address, options.Database, options.Charset, options.ParseTime, options.Location)
	}

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
