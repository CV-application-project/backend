package service

import (
	"Backend-Server/cv_service/api"
	"Backend-Server/cv_service/config"
	"Backend-Server/cv_service/store"
	"context"
	"github.com/go-logr/logr"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Service struct {
	log   logr.Logger
	cfg   *config.Config
	store store.Querier
	api.UnimplementedCVServiceServer
}

func (s *Service) RegisterWithServer(server *grpc.Server) {
	api.RegisterCVServiceServer(server, s)
}

func (s *Service) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := api.RegisterCVServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}
	return err
}

func NewService(logger logr.Logger, store store.Querier, cfg *config.Config) *Service {
	return &Service{
		cfg:   cfg,
		log:   logger,
		store: store,
	}
}

func (s *Service) Close(_ context.Context) {
	err := s.store.Close()
	if err != nil {
		s.log.Info("can't not close db")
	}
}
