package server

import "net/http"

type HTTPServerHandler func(http.Handler) http.Handler
type HTTPServerMiddleware func(http.Handler) http.Handler

func WithGrpcAddrListen(l Listen) Option {
	return func(c *Config) {
		c.Grpc.Addr = l
	}
}

func WithServiceServer(srv ...ServiceServer) Option {
	return func(c *Config) {
		c.ServiceServers = append(c.ServiceServers, srv...)
	}
}

func WithHttpAddrListen(l Listen) Option {
	return func(c *Config) {
		c.HttpGateway.Addr = l
	}
}

func WithGatewayServerHandler(handlers ...HTTPServerHandler) Option {
	return func(c *Config) {
		c.HttpGateway.ServerHandlers = append(c.HttpGateway.ServerHandlers, handlers...)
	}
}
