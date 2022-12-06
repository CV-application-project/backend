package service

import (
	"Backend-Server/gateway/api"
	"Backend-Server/gateway/constant"
	"Backend-Server/gateway/store"
	userApi "Backend-Server/user_service/api"
	"context"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (s *Service) RegisterNewUser(ctx context.Context, req *api.RegisterNewUserRequest) (*api.RegisterNewUserResponse, error) {
	logger := s.log.WithName("RegisterNewUser")
	if err := req.Validate(); err != nil {
		logger.Error(err, "Validate request failed")
		return nil, err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), constant.PasswordCost)
	if err != nil {
		logger.Error(err, "bcrypt | GenerateFromPassword | Can not hash password")
		return nil, err
	}

	userReq := &userApi.RegisterUserRequest{
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashPassword),
	}

	// Send request to User service
	res, err := s.userClient.RegisterUser(ctx, userReq)
	if err != nil {
		logger.Error(err, "userClient | RegisterUser | Error from user client while register new user")
		return nil, err
	}
	// Create Bearer token
	userToken, err := createBearerToken(res.Data)
	if err != nil {
		logger.Error(err, "createBearerToken | can not create Bearer token")
		return nil, err
	}
	// Store token into db
	if _, err = s.store.CreateUserInfo(ctx, store.CreateUserInfoParams{
		UserID:    res.Data.UserId,
		Username:  req.Username,
		Email:     req.Email,
		Token:     userToken.Token,
		ExpiredAt: userToken.ExpiredAt,
	}); err != nil {
		logger.Error(err, "Store | CreateUserInfo | Can not create new user info", "user_id", res.Data.UserId)
		return nil, err
	}

	return &api.RegisterNewUserResponse{
		Code:    http.StatusOK,
		Message: "Register new user successfully",
		Token:   "Bearer " + userToken.Token,
	}, nil
}

func (s *Service) AuthorizeUser(ctx context.Context, req *api.AuthorizeUserRequest) (*api.AuthorizeUserResponse, error) {
	logger := s.log.WithName("AuthorizeUser")
	if err := req.Validate(); err != nil {
		logger.Error(err, "Validate request failed")
		return nil, err
	}
	return &api.AuthorizeUserResponse{
		Code:    http.StatusOK,
		Message: "Authorize user successfully",
		Token:   "Chua co dau",
	}, nil
}
