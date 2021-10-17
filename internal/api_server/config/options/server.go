package options

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

const (
	FlagSetName       = "server"
	FlagServerPort    = "server.port"
	FlagServerMode    = "server.mode"
	FlagServerAddress = "server.address"
)

type ServerOptions struct {
	Address string `json:"address" yaml:"address"`
	Port    int    `json:"port" yaml:"port"`
	Mode    string `json:"mode" yaml:"mode"`
	//Middlewares []string `json:"middlewares"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Address: "0.0.0.0",
		Port:    8080,
		Mode:    gin.ReleaseMode,
		//Middlewares: []string{},
	}
}

func (s *ServerOptions) AddFlag() *pflag.FlagSet {
	fs := pflag.NewFlagSet(FlagSetName, pflag.ExitOnError)

	fs.IntVar(&s.Port, FlagServerPort, s.Port, "server port")
	fs.StringVar(&s.Address, FlagServerAddress, s.Address, "server address")
	fs.StringVar(&s.Mode, FlagServerMode, s.Mode, "server mode, supported: debug/test/release.")

	return fs
}
