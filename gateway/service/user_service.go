package service

import (
	"Backend-Server/common/ctx_key"
	"Backend-Server/gateway/api"
	"Backend-Server/gateway/store"
	userApi "Backend-Server/user_service/api"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"time"
)

const TimeOneDay = 24 * time.Hour

func (s *Service) RegisterNewUser(ctx context.Context, req *api.RegisterNewUserRequest) (*api.RegisterNewUserResponse, error) {
	logger := s.log.WithName("RegisterNewUser")
	if err := req.Validate(); err != nil {
		logger.Error(err, "Validate request failed")
		return nil, err
	}

	userReq := &userApi.RegisterUserRequest{
		EmployeeId: req.EmployeeId,
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		Address:    req.Address,
		Phone:      req.Phone,
		Gender:     req.Gender,
		Department: req.Department,
		Position:   req.Position,
		Role:       toUserRole(req.Role),
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
		UserID:     res.Data.UserId,
		EmployeeID: req.EmployeeId,
		Email:      req.Email,
		Token:      userToken.Token,
		ExpiredAt:  userToken.ExpiredAt,
		Role:       sql.NullString{Valid: true, String: req.Role.String()},
		Department: sql.NullString{Valid: true, String: req.Department},
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

	user, err := s.store.GetUserInfoByUsernameOrEmail(ctx, store.GetUserInfoByUsernameOrEmailParams{
		EmployeeID: req.EmployeeId,
		Email:      req.Email,
	})
	if err != nil {
		logger.Error(err, "Store | GetUserInfoByUsernameOrEmail")
		return nil, err
	}

	authorizeResp, err := s.userClient.AuthorizeUser(ctx, &userApi.AuthorizeUserRequest{
		EmployeeId: user.EmployeeID,
		Email:      user.Email,
		Password:   req.Password,
	})
	if err != nil {
		logger.Error(err, "userClient | AuthorizeUser")
		return nil, err
	}

	if user.ExpiredAt.After(time.Now().Add(TimeOneDay)) {
		return &api.AuthorizeUserResponse{
			Code:    http.StatusOK,
			Message: "success",
			Token:   "Bearer " + user.Token,
		}, nil
	}

	userToken, err := createBearerToken(authorizeResp.Data)
	if err != nil {
		logger.Error(err, "createBearerToken | can not create Bearer token")
		return nil, err
	}

	if _, err = s.store.UpdateUserInfoTokenByUserId(ctx, store.UpdateUserInfoTokenByUserIdParams{
		UserID:    user.UserID,
		Token:     userToken.Token,
		ExpiredAt: userToken.ExpiredAt,
	}); err != nil {
		logger.Error(err, "Store | UpdateUserInfoTokenByUserId")
		return nil, err
	}

	return &api.AuthorizeUserResponse{
		Code:    http.StatusOK,
		Message: "success",
		Token:   "Bearer " + userToken.Token,
	}, nil
}

func (s *Service) HTTPUpdateUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "PUT" {
		err := errors.New("invalid HTTP method")
		return err
	}
	var info userApi.UpdateInfo
	data := bytes.NewBuffer(nil)
	if _, err := io.Copy(data, r.Body); err != nil {
		return err
	}
	if err := json.Unmarshal(data.Bytes(), &info); err != nil {
		return err
	}
	resp, err := s.userClient.UpdateUserInfo(context.Background(), &userApi.UpdateUserInfoRequest{
		AssignerId:   cast.ToInt64(r.Context().Value(ctx_key.ContextUserId)),
		AssignerRole: convertStringToRole(r.Context().Value(ctx_key.ContextRole)),
		UpdateData:   &info,
	})
	if err != nil {
		return err
	}
	if _, err = s.store.UpdateUserCard(context.Background(), store.UpdateUserCardParams{
		FrontCard: sql.NullString{Valid: true, String: resp.Data.FrontCard},
		BackCard:  sql.NullString{Valid: true, String: resp.Data.BackCard},
	}); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
	return nil
}

func convertStringToRole(value any) userApi.UserRole {
	switch value {
	case "STAFF":
		return userApi.UserRole_STAFF
	case "MANAGER":
		return userApi.UserRole_MANAGER
	case "HR":
		return userApi.UserRole_HR
	default:
		return userApi.UserRole_OTHER
	}
}

func (s *Service) HTTPExportUsersByDepartment(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return errors.New("invalid http method")
	}
	resp, err := s.userClient.ExportUsersByDepartment(context.Background(), &userApi.ExportUsersByDepartmentRequest{
		AssignerId:         cast.ToInt64(r.Context().Value(ctx_key.ContextUserId)),
		AssignerRole:       convertStringToRole(r.Context().Value(ctx_key.ContextRole)),
		AssignerDepartment: cast.ToString(r.Context().Value(ctx_key.ContextDepartment)),
	})
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
	return nil
}
