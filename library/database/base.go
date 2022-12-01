package database

import "Backend-Server/library/server"

type Base struct {
	Env    string       `json:"env" mapstructure:"env"`
	Server ServerConfig `json:"server" mapstructure:"server"`
}

type ServerConfig struct {
	GRPC server.Listen `json:"grpc" mapstructure:"grpc" yaml:"grpc"`
	HTTP server.Listen `json:"http" mapstructure:"http" yaml:"http"`
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		GRPC: server.Listen{
			Host: "0.0.0.0",
			Port: 10443,
		},
		HTTP: server.Listen{
			Host: "0.0.0.0",
			Port: 10080,
		},
	}
}

func DefaultBaseConfig() *Base {
	return &Base{
		Env:    "env",
		Server: DefaultServerConfig(),
	}
}
