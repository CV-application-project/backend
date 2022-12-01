package server

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

const (
	SERVER_OPTION_MAX_CONNECTION_AGE       = 120 // seconds
	SERVER_OPTION_MAX_CONNECTION_AGE_GRACE = 150 // seconds
	SERVER_OPTION_MAX_CONNECTION_IDLE      = 30  // seconds
	SERVER_OPTION_ENFORCEMENT_MIN_TIME     = 5   // seconds
)

type grpcConfig struct {
	Addr                               Listen
	PreServerUnaryInterceptors         []grpc.UnaryServerInterceptor
	PreServerStreamInterceptors        []grpc.StreamServerInterceptor
	PostServerUnaryInterceptors        []grpc.UnaryServerInterceptor
	PostServerStreamInterceptors       []grpc.StreamServerInterceptor
	DefaultUnaryInterceptorsValidator  grpc.UnaryServerInterceptor
	DefaultStreamInterceptorsValidator grpc.StreamServerInterceptor
	ServerOption                       []grpc.ServerOption
	MaxConcurrentStreams               uint32
}

func createDefaultGrpcConfig() *grpcConfig {
	grpc_prometheus.EnableHandlingTimeHistogram()
	return &grpcConfig{
		Addr: Listen{
			Host: "0.0.0.0",
			Port: 10443,
		},
		PreServerUnaryInterceptors: []grpc.UnaryServerInterceptor{
			otelgrpc.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		},
		PreServerStreamInterceptors: []grpc.StreamServerInterceptor{
			otelgrpc.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		},
		DefaultUnaryInterceptorsValidator:  grpc_validator.UnaryServerInterceptor(),
		DefaultStreamInterceptorsValidator: grpc_validator.StreamServerInterceptor(),
		MaxConcurrentStreams:               1000,
	}
}

type grpcServer struct {
	server *grpc.Server
	config *grpcConfig
}

func (c *grpcConfig) buildServerUnaryInterceptors() []grpc.UnaryServerInterceptor {
	unaryInterceptors := c.PreServerUnaryInterceptors
	unaryInterceptors = append(
		unaryInterceptors,
		c.DefaultUnaryInterceptorsValidator,
	)
	unaryInterceptors = append(unaryInterceptors, c.PostServerUnaryInterceptors...)
	return unaryInterceptors
}

func (c *grpcConfig) buildServerStreamInterceptors() []grpc.StreamServerInterceptor {
	streamInterceptors := c.PreServerStreamInterceptors
	streamInterceptors = append(
		streamInterceptors,
		c.DefaultStreamInterceptorsValidator,
	)
	streamInterceptors = append(streamInterceptors, c.PostServerStreamInterceptors...)
	return streamInterceptors
}

func (c *grpcConfig) ServerOptions() []grpc.ServerOption {
	enforcement := keepalive.EnforcementPolicy{
		MinTime:             SERVER_OPTION_ENFORCEMENT_MIN_TIME * time.Second,
		PermitWithoutStream: true,
	}
	return append(
		[]grpc.ServerOption{
			grpc_middleware.WithUnaryServerChain(c.buildServerUnaryInterceptors()...),
			grpc_middleware.WithStreamServerChain(c.buildServerStreamInterceptors()...),
			grpc.MaxConcurrentStreams(c.MaxConcurrentStreams),
			grpc.KeepaliveEnforcementPolicy(enforcement), // here
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle:     SERVER_OPTION_MAX_CONNECTION_IDLE * time.Second,
				MaxConnectionAge:      SERVER_OPTION_MAX_CONNECTION_AGE * time.Second,
				MaxConnectionAgeGrace: SERVER_OPTION_MAX_CONNECTION_AGE_GRACE * time.Second,
			}),
		},
		c.ServerOption...,
	)
}

func newGrpcServer(c *grpcConfig, servers []ServiceServer) *grpcServer {
	s := grpc.NewServer(c.ServerOptions()...)
	for _, svr := range servers {
		svr.RegisterWithServer(s)
	}
	return &grpcServer{
		server: s,
		config: c,
	}
}

func (s *grpcServer) Serve() error {
	l, err := s.config.Addr.CreateListener()
	if err != nil {
		return fmt.Errorf("failed to create listener %w", err)
	}

	log.Println("gRPC server is starting ", l.Addr())

	err = s.server.Serve(l)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to serve gRPC server %w", err)
	}
	log.Println("gRPC server ready")

	return nil
}

func (s *grpcServer) Shutdown(ctx context.Context) {
	s.server.GracefulStop()
}
