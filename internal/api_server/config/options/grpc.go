package options

import (
	"github.com/spf13/pflag"
)

type GRPCOptions struct {
	Address    string `json:"address"`
	Port       int    `json:"port"`
	MaxMsgSize int    `json:"max-msg-size"`
}

func NewGRPCOptions() *GRPCOptions {
	return &GRPCOptions{
		Address:    "0.0.0.0",
		Port:       8081,
		MaxMsgSize: 4 * 1024 * 1024,
	}
}

func (s *GRPCOptions) AddFlag() *pflag.FlagSet {
	fs := pflag.NewFlagSet("grpc", pflag.ExitOnError)

	fs.StringVar(&s.Address, "grpc.address", s.Address, ""+
		"The IP address on which to serve the --grpc.bind-port(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")

	fs.IntVar(&s.Port, "grpc.port", s.Port, ""+
		"The port on which to serve unsecured, unauthenticated grpc access. It is assumed "+
		"that firewall rules are set up such that this port is not reachable from outside of "+
		"the deployed machine and that port 443 on the iam public address is proxied to this "+
		"port. This is performed by nginx in the default setup. Set to zero to disable.")

	fs.IntVar(&s.MaxMsgSize, "grpc.max-msg-size", s.MaxMsgSize, "gRPC max message size.")

	return fs
}
