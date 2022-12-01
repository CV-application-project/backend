package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	grpcServer *grpcServer
	httpServer *httpServer
	config     *Config
}

func New(opts ...Option) (*Server, error) {
	cfg := createConfig(opts)

	log.Println("Create grpc server")
	grpcServer := newGrpcServer(cfg.Grpc, cfg.ServiceServers)
	reflection.Register(grpcServer.server)

	conn, err := grpc.Dial(cfg.Grpc.Addr.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("fail to dial gRPC server. %w", err)
	}

	log.Println("Create gateway server")
	httpServer, err := newGatewayServer(cfg.HttpGateway, conn, cfg.ServiceServers)
	if err != nil {
		return nil, fmt.Errorf("fail to create gateway server. %w", err)
	}
	return &Server{
		grpcServer: grpcServer,
		httpServer: httpServer,
		config:     cfg,
	}, nil
}

func (s *Server) Serve() error {
	stop := make(chan os.Signal, 1)
	channelErr := make(chan error)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := s.httpServer.Serve(); err != nil {
			log.Println("Error starting http server, ", err)
			channelErr <- err
		}
	}()

	go func() {
		if err := s.grpcServer.Serve(); err != nil {
			log.Println("Error starting gRPC server, ", err)
			channelErr <- err
		}
	}()

	// shutdown
	for {
		select {
		//goland:noinspection GoDeferInLoop
		case <-stop:
			log.Println("Shutting down server")

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			for _, ss := range s.config.ServiceServers {
				ss.Close(ctx)
			}

			s.httpServer.Shutdown(ctx)
			s.grpcServer.Shutdown(ctx)
			return nil
		case err := <-channelErr:
			return err
		}
	}
}
