package client

import (
	"Backend-Server/timekeeping_service/api"
	"context"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TimekeepingClient interface {
	GetHistoryOfUser(ctx context.Context, req *api.GetHistoryOfUserRequest) (*api.GetHistoryOfUserResponse, error)
	CreateHistoryOfUser(ctx context.Context, req *api.CreateHistoryOfUserRequest) (*api.CreateHistoryOfUserResponse, error)
	UpdateHistoryOfUser(ctx context.Context, req *api.UpdateHistoryOfUserRequest) (*api.UpdateHistoryOfUserResponse, error)
}

type timekeepingClient struct {
	client api.TimekeepingServiceClient
	log    logr.Logger
}

func NewTimekeepingClient(logger logr.Logger, addr string) (TimekeepingClient, error) {
	clientConn, err := grpc.DialContext(context.Background(), addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &timekeepingClient{
		client: api.NewTimekeepingServiceClient(clientConn),
		log:    logger,
	}, nil
}

func (c *timekeepingClient) GetHistoryOfUser(ctx context.Context, req *api.GetHistoryOfUserRequest) (*api.GetHistoryOfUserResponse, error) {
	return c.client.GetHistoryOfUser(ctx, req)
}

func (c *timekeepingClient) CreateHistoryOfUser(ctx context.Context, req *api.CreateHistoryOfUserRequest) (*api.CreateHistoryOfUserResponse, error) {
	return c.client.CreateHistoryOfUser(ctx, req)
}

func (c *timekeepingClient) UpdateHistoryOfUser(ctx context.Context, req *api.UpdateHistoryOfUserRequest) (*api.UpdateHistoryOfUserResponse, error) {
	return c.client.UpdateHistoryOfUser(ctx, req)
}
