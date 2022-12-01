package service

import (
	"Backend-Server/gateway/api"
	"context"
	"fmt"
	"net/http"
)

func (s *Service) RegisterNewUser(ctx context.Context, req *api.RegisterNewUserRequest) (*api.RegisterNewUserResponse, error) {
	fmt.Printf("Req: %+v\n", req)
	return &api.RegisterNewUserResponse{
		Code:    http.StatusOK,
		Message: "Hello from gateway service",
	}, nil
}

func (s *Service) HelloWorld(ctx context.Context, req *api.HelloWorldRequest) (*api.HelloWorldResponse, error) {
	return &api.HelloWorldResponse{
		Message: "Welcome",
	}, nil
}
