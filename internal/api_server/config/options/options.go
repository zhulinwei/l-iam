package options

import "github.com/spf13/pflag"

type IOptions interface {
	AddFlag() *pflag.FlagSet
}

// Options 配置：用来构建命令行参数，它的值来自于命令行选项或者配置文件
type Options struct {
	Jwt    *JwtOptions    `json:"jwt"`
	Grpc   *GRPCOptions   `json:"grpc"`
	MySQL  *MySQLOptions  `json:"mysql"`
	Redis  *RedisOptions  `json:"redis"`
	Server *ServerOptions `json:"server"`
}

func NewOptions() *Options {
	return &Options{
		Jwt:    NewJwtOptions(),
		Grpc:   NewGRPCOptions(),
		MySQL:  NewMySQLOptions(),
		Redis:  NewRedisOptions(),
		Server: NewServerOptions(),
	}
}

func (o *Options) Flags() []*pflag.FlagSet {
	var flags []*pflag.FlagSet
	flags = append(flags, o.Jwt.AddFlag())
	flags = append(flags, o.Grpc.AddFlag())
	flags = append(flags, o.MySQL.AddFlag())
	flags = append(flags, o.Redis.AddFlag())
	flags = append(flags, o.Server.AddFlag())

	return flags
}
