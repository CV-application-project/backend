package errorz

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound           = status.Error(codes.FailedPrecondition, "User not exists in database")
	ErrAssignerRoleNotAllowed = status.Error(codes.PermissionDenied, "Assigner role not allowed")
)
