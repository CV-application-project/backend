package service

import (
	"Backend-Server/gateway/api"
	userApi "Backend-Server/user_service/api"
)

func toUserRole(role api.UserRole) userApi.UserRole {
	switch role {
	case api.UserRole_HR:
		return userApi.UserRole_HR
	case api.UserRole_MANAGER:
		return userApi.UserRole_MANAGER
	case api.UserRole_STAFF:
		return userApi.UserRole_STAFF
	default:
		return userApi.UserRole_OTHER
	}
}
