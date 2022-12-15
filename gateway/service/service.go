package service

import (
	"Backend-Server/gateway/api"
	"Backend-Server/gateway/client"
	"Backend-Server/gateway/config"
	"Backend-Server/gateway/store"
	"context"
	"github.com/go-logr/logr"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Service struct {
	log        logr.Logger
	cfg        *config.Config
	store      store.Querier
	userClient client.UserClient
	cvClient   client.CVClient
	api.UnimplementedGatewayServiceServer
}

func (s *Service) RegisterWithServer(server *grpc.Server) {
	api.RegisterGatewayServiceServer(server, s)
}

func (s *Service) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := api.RegisterGatewayServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}
	return err
}

func NewService(logger logr.Logger, store store.Querier, cfg *config.Config) *Service {
	userClient, err := client.NewUserClient(logger, cfg.ClientHost.UserService)
	if err != nil {
		userClient = nil
	}
	cvClient, err := client.NewCVClient(logger, cfg.ClientHost.CVService)
	if err != nil {
		cvClient = nil
	}
	return &Service{
		cfg:        cfg,
		log:        logger,
		store:      store,
		userClient: userClient,
		cvClient:   cvClient,
	}
}

func (s *Service) Close(_ context.Context) {
	err := s.store.Close()
	if err != nil {
		s.log.Info("can't not close db")
	}
}
