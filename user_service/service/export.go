package service

import (
	"Backend-Server/user_service/api"
	"Backend-Server/user_service/store"
	"context"
	"database/sql"
	"net/http"
)

func (s *Service) ExportUsersByDepartment(ctx context.Context, req *api.ExportUsersByDepartmentRequest) (*api.ExportUsersByDepartmentResponse, error) {
	if _, err := s.store.GetUserInfoById(ctx, req.AssignerId); err != nil {
		return nil, err
	}
	var users []store.User
	var err error
	switch req.AssignerRole {
	case api.UserRole_HR:
		users, err = s.store.GetAllUsers(ctx)
		if err != nil {
			return nil, err
		}
	case api.UserRole_MANAGER:
		users, err = s.store.GetUsersByDepartment(ctx, sql.NullString{Valid: true, String: req.AssignerDepartment})
		if err != nil {
			return nil, err
		}
	default:
		return nil, nil
	}
	return &api.ExportUsersByDepartmentResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    toApiUser(users),
	}, nil
}

func toApiUser(users []store.User) []*api.User {
	if len(users) == 0 {
		return nil
	}
	data := make([]*api.User, 0, len(users))
	for _, user := range users {
		data = append(data, &api.User{
			Position:   user.Position.String,
			Role:       user.Role.String,
			Address:    user.Address.String,
			Department: user.Department.String,
			Phone:      user.Phone.String,
			Gender:     user.Gender.String,
			Name:       user.Name,
			Email:      user.Email,
			EmployeeId: user.EmployeeID,
			FrontCard:  user.FrontCard.String,
			BackCard:   user.BackCard.String,
		})
	}
	return data
}
