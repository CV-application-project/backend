package service

import (
	"Backend-Server/user_service/api"
	"Backend-Server/user_service/constant"
	"Backend-Server/user_service/store"
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (s *Service) RegisterUser(ctx context.Context, req *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	logger := s.log.WithName("RegisterUser").WithValues("traceId", req.Username)
	if err := req.Validate(); err != nil {
		logger.Error(err, "Validate request failed")
		return nil, err
	}
	// Check whether user with this username exists?
	_, err := s.store.GetUserByUsernameOrEmail(ctx, store.GetUserByUsernameOrEmailParams{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err, "GetUserByUsernameOrEmail | Can not get user record", "username", req.Username, "email", req.Email)
			return nil, err
		}
		logger.Info("There is no user with this username and email")
		// Case user doesn't exist
		insertResult, err := s.store.CreateNewUserInfo(ctx, store.CreateNewUserInfoParams{
			Name:     req.Name,
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
			Data:     sql.NullString{Valid: false, String: ""},
		})
		if err != nil {
			logger.Error(err, "CreateNewUserInfo | Can not create new user", "username", req.Username)
			return nil, err
		}
		userId, err := insertResult.LastInsertId()
		if err != nil {
			logger.Error(err, "Store | LastInsertId | Can not get user id")
			return nil, err
		}

		// Create new user token
		tokenParam := store.CreateNewUserTokenByUserIdParams{
			UserID:    userId,
			Token:     uuid.New().String(),
			ExpiredAt: time.Now().UTC().Add(constant.TokenExpiredDay * 24 * time.Hour),
		}
		if _, err = s.store.CreateNewUserTokenByUserId(ctx, tokenParam); err != nil {
			logger.Error(err, "Store | CreateNewUserTokenByUserId | Can not create new user token", "user id", userId)
			return nil, err
		}

		return &api.RegisterUserResponse{
			Code:    http.StatusOK,
			Message: "Register new user successfully",
			Data: &api.UserToken{
				UserId: tokenParam.UserID,
				Token:  tokenParam.Token,
				ExpiredAt: &timestamp.Timestamp{
					Seconds: tokenParam.ExpiredAt.Unix(),
					Nanos:   0,
				},
			},
		}, nil
	}
	// case user already exists
	return &api.RegisterUserResponse{
		Code:    http.StatusOK,
		Message: "User already exists",
		Data:    nil,
	}, nil
}
