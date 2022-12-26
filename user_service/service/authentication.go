package service

import (
	"Backend-Server/user_service/api"
	"Backend-Server/user_service/constant"
	"Backend-Server/user_service/store"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (s *Service) RegisterUser(ctx context.Context, req *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	logger := s.log.WithName("RegisterUser").WithValues("traceId", req.EmployeeId)
	// Check whether user with this username exists?
	_, err := s.store.GetUserByUsernameOrEmail(ctx, store.GetUserByUsernameOrEmailParams{
		Username: req.EmployeeId,
		Email:    req.Email,
	})
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err, "GetUserByUsernameOrEmail | Can not get user record", "username", req.EmployeeId, "email", req.Email)
			return nil, err
		}
		logger.Info("There is no user with this username and email")
		// Case user doesn't exist

		hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), constant.PasswordCost)
		if err != nil {
			logger.Error(err, "bcrypt | GenerateFromPassword | Can not hash password")
			return nil, err
		}

		insertResult, err := s.store.CreateNewUserInfo(ctx, store.CreateNewUserInfoParams{
			Name:       req.Name,
			Username:   req.EmployeeId,
			Password:   string(hashPassword),
			Email:      req.Email,
			Role:       sql.NullString{Valid: true, String: req.Role.String()},
			Position:   sql.NullString{Valid: true, String: req.Position},
			Department: sql.NullString{Valid: true, String: req.Department},
			Phone:      sql.NullString{Valid: true, String: req.Phone},
			Address:    sql.NullString{Valid: true, String: req.Address},
			Gender:     sql.NullString{Valid: true, String: req.Gender},
			Data:       sql.NullString{Valid: false, String: ""},
		})
		if err != nil {
			logger.Error(err, "CreateNewUserInfo | Can not create new user", "username", req.EmployeeId)
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
				UserId:    tokenParam.UserID,
				Token:     tokenParam.Token,
				ExpiredAt: tokenParam.ExpiredAt.Unix(),
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

func (s *Service) AuthorizeUser(ctx context.Context, req *api.AuthorizeUserRequest) (*api.AuthorizeUserResponse, error) {
	traceId := ""
	if req.EmployeeId != "" {
		traceId = req.EmployeeId
	} else if req.Email != "" {
		traceId = req.Email
	}
	logger := s.log.WithName("AuthorizeUser").WithValues("userId", traceId)

	user, err := s.store.GetUserByUsernameOrEmail(ctx, store.GetUserByUsernameOrEmailParams{
		Username: req.EmployeeId,
		Email:    req.Email,
	})
	if err != nil {
		logger.Error(err, "Store | GetUserByUsernameOrEmail")
		return nil, err
	}

	// Validate password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		logger.Error(err, "bcrypt | CompareHashAndPassword")
		return nil, err
	}

	token, err := s.store.GetUserTokenByUserId(ctx, user.ID)
	if err != nil {
		logger.Error(err, "Store | GetUserTokenByUserId")
		return nil, err
	}
	return &api.AuthorizeUserResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data: &api.UserToken{
			UserId:    token.UserID,
			Token:     token.Token,
			ExpiredAt: token.ExpiredAt.Unix(),
		},
	}, nil
}
