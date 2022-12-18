package client

import (
	cvApi "Backend-Server/cv_service/api"
	"context"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CVClient interface {
	RegisterCICForUser(ctx context.Context, req *cvApi.RegisterCICForUserRequest) (*cvApi.RegisterCICForUserResponse, error)
	RegisterUserFace(ctx context.Context, req *cvApi.RegisterUserFaceRequest) (*cvApi.RegisterUserFaceResponse, error)
	AuthorizeUserFace(ctx context.Context, req *cvApi.AuthorizeUserFaceRequest) (*cvApi.AuthorizeUserFaceResponse, error)
}

type cvClient struct {
	client cvApi.CVServiceClient
	log    logr.Logger
}

func NewCVClient(logger logr.Logger, addr string) (CVClient, error) {
	clientConn, err := grpc.DialContext(context.Background(), addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &cvClient{
		client: cvApi.NewCVServiceClient(clientConn),
		log:    logger,
	}, nil
}

func (c *cvClient) RegisterCICForUser(ctx context.Context, req *cvApi.RegisterCICForUserRequest) (*cvApi.RegisterCICForUserResponse, error) {
	return c.client.RegisterCICForUser(ctx, req)
}

func (c *cvClient) RegisterUserFace(ctx context.Context, req *cvApi.RegisterUserFaceRequest) (*cvApi.RegisterUserFaceResponse, error) {
	return c.client.RegisterUserFace(ctx, req)
}

func (c *cvClient) AuthorizeUserFace(ctx context.Context, req *cvApi.AuthorizeUserFaceRequest) (*cvApi.AuthorizeUserFaceResponse, error) {
	return c.client.AuthorizeUserFace(ctx, req)
}
