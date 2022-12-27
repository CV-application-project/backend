package common_error

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound = status.Error(codes.FailedPrecondition, "User not exists in database")
)
