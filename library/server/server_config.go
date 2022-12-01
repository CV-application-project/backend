package server

import (
	"fmt"
	"net"
)

type Listen struct {
	Host string `json:"host" mapstructure:"host" yaml:"host"`
	Port int    `json:"port" mapstructure:"port" yaml:"port"`
}

func (l *Listen) String() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

func (l *Listen) CreateListener() (net.Listener, error) {
	lis, err := net.Listen("tcp", l.String())
	if err != nil {
		return nil, fmt.Errorf("failed to listen %s: %w", l.String(), err)
	}
	return lis, nil
}

type Config struct {
	HttpGateway    *httpConfig
	Grpc           *grpcConfig
	ServiceServers []ServiceServer
}

func createDefaultConfig() *Config {
	return &Config{
		Grpc:        createDefaultGrpcConfig(),
		HttpGateway: createDefaultHttpConfig(),
	}
}

type Option func(*Config)

func createConfig(opts []Option) *Config {
	c := createDefaultConfig()
	for _, f := range opts {
		f(c)
	}
	return c
}
