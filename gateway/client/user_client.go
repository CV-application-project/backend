package client

import (
	"Backend-Server/user_service/api"
	"context"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient interface {
	RegisterUser(ctx context.Context, req *api.RegisterUserRequest) (*api.RegisterUserResponse, error)
	AuthorizeUser(ctx context.Context, req *api.AuthorizeUserRequest) (*api.AuthorizeUserResponse, error)
}

type userClient struct {
	client api.UserServiceClient
	log    logr.Logger
}

func NewUserClient(logger logr.Logger, addr string) (UserClient, error) {
	clientConn, err := grpc.DialContext(context.Background(), addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &userClient{
		client: api.NewUserServiceClient(clientConn),
		log:    logger,
	}, nil
}

func (c *userClient) RegisterUser(ctx context.Context, req *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	return c.client.RegisterUser(ctx, req)
}

func (c *userClient) AuthorizeUser(ctx context.Context, req *api.AuthorizeUserRequest) (*api.AuthorizeUserResponse, error) {
	return c.client.AuthorizeUser(ctx, req)
}
