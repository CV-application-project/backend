// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package store

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateNewUserInfo(ctx context.Context, arg CreateNewUserInfoParams) (sql.Result, error)
	CreateNewUserTokenByUserId(ctx context.Context, arg CreateNewUserTokenByUserIdParams) (sql.Result, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserByUsernameOrEmail(ctx context.Context, arg GetUserByUsernameOrEmailParams) (User, error)
	GetUserInfoById(ctx context.Context, id int64) (User, error)
	GetUserTokenByUserId(ctx context.Context, userID int64) (UserToken, error)
	GetUsersByDepartment(ctx context.Context, department sql.NullString) ([]User, error)
	UpdateUserInfoById(ctx context.Context, arg UpdateUserInfoByIdParams) (sql.Result, error)
	Close() error
}

var _ Querier = (*Queries)(nil)
