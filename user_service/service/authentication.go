package service

import (
	"Backend-Server/common/errorz"
	"Backend-Server/common/helper"
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
	if req.Email == helper.EmptyString {
		req.Email = uuid.New().String()
	}
	user, err := s.store.GetUserByUsernameOrEmail(ctx, store.GetUserByUsernameOrEmailParams{
		EmployeeID: req.EmployeeId,
		Email:      req.Email,
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
			EmployeeID: req.EmployeeId,
			Password:   string(hashPassword),
			Email:      req.Email,
			Role:       sql.NullString{Valid: true, String: req.Role.String()},
			Position:   sql.NullString{Valid: true, String: req.Position},
			Department: sql.NullString{Valid: true, String: req.Department},
			Phone:      sql.NullString{Valid: true, String: req.Phone},
			Address:    sql.NullString{Valid: true, String: req.Address},
			Gender:     sql.NullString{Valid: true, String: req.Gender},
			Data:       sql.NullString{Valid: false, String: ""},
			FrontCard:  sql.NullString{Valid: true},
			BackCard:   sql.NullString{Valid: true},
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
	info, err := s.store.GetUserTokenByUserId(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &api.RegisterUserResponse{
		Code:    http.StatusOK,
		Message: "User already exists",
		Data: &api.UserToken{
			UserId:    info.UserID,
			Token:     info.Token,
			ExpiredAt: info.ExpiredAt.Unix(),
		},
	}, nil
}

func (s *Service) AuthorizeUser(ctx context.Context, req *api.AuthorizeUserRequest) (*api.AuthorizeUserResponse, error) {
	traceId := ""
	if req.EmployeeId != helper.EmptyString {
		traceId = req.EmployeeId
	} else if req.Email != helper.EmptyString {
		traceId = req.Email
	}
	logger := s.log.WithName("AuthorizeUser").WithValues("userId", traceId)
	if req.Email == helper.EmptyString {
		req.Email = uuid.New().String()
	}
	user, err := s.store.GetUserByUsernameOrEmail(ctx, store.GetUserByUsernameOrEmailParams{
		EmployeeID: req.EmployeeId,
		Email:      req.Email,
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

func (s *Service) UpdateUserInfo(ctx context.Context, req *api.UpdateUserInfoRequest) (*api.UpdateUserInfoResponse, error) {
	logger := s.log.WithName("UpdateUserInfo").WithValues("assigner", req.AssignerId, "employee", req.UpdateData.UserId)
	user, err := s.store.GetUserInfoById(ctx, req.UpdateData.UserId)
	if err != nil {
		logger.Error(err, "store | GetUserInfoById")
		return nil, err
	}
	switch req.AssignerRole {
	case api.UserRole_HR:
		if req.UpdateData.Department != helper.EmptyString {
			user.Department.String = req.UpdateData.Department
		}
		if req.UpdateData.Role.String() != helper.EmptyString {
			user.Role.String = req.UpdateData.Role.String()
		}
		if req.UpdateData.Position != helper.EmptyString {
			user.Position.String = req.UpdateData.Position
		}
	case api.UserRole_OTHER,
		api.UserRole_MANAGER:
		return nil, errorz.ErrAssignerRoleNotAllowed
	default:
		break
	}
	if req.UpdateData.Address != helper.EmptyString {
		user.Address.String = req.UpdateData.Address
	}
	if req.UpdateData.Phone != helper.EmptyString {
		user.Phone.String = req.UpdateData.Phone
	}
	if req.UpdateData.FrontCard != helper.EmptyString {
		user.FrontCard.String = req.UpdateData.FrontCard
	}
	if req.UpdateData.BackCard != helper.EmptyString {
		user.BackCard.String = req.UpdateData.BackCard
	}
	if _, err = s.store.UpdateUserInfoById(ctx, store.UpdateUserInfoByIdParams{
		Phone:      user.Phone,
		Department: user.Department,
		Address:    user.Address,
		Role:       user.Role,
		ID:         user.ID,
		Position:   user.Position,
		FrontCard:  user.FrontCard,
		BackCard:   user.BackCard,
	}); err != nil {
		logger.Error(err, "store | UpdateUserInfoById")
		return nil, err
	}
	return &api.UpdateUserInfoResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    toApiUser([]store.User{user})[0],
	}, nil
}
