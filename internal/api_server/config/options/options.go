package options

import "github.com/spf13/pflag"

type IOptions interface {
	AddFlag() *pflag.FlagSet
}

type Options struct {
	//Jwt    *JwtOptions
	//Grpc   *GRPCOptions
	//MySQL  *MySQLOptions
	//Redis  *RedisOptions
	Server *ServerOptions
}

func NewOptions() *Options {
	return &Options{
		//Jwt:    NewJwtOptions(),
		//Grpc:   NewGRPCOptions(),
		//MySQL:  NewMySQLOptions(),
		//Redis:  NewRedisOptions(),
		Server: NewServerOptions(),
	}
}

func (o *Options) Flags() []*pflag.FlagSet {
	var flags []*pflag.FlagSet
	//flags = append(flags, o.Jwt.AddFlag())
	//flags = append(flags, o.Grpc.AddFlag())
	//flags = append(flags, o.MySQL.AddFlag())
	//flags = append(flags, o.Redis.AddFlag())
	flags = append(flags, o.Server.AddFlag())

	return flags
}
