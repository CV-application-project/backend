// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user_info.sql

package store

import (
	"context"
	"database/sql"
)

const createNewUserInfo = `-- name: CreateNewUserInfo :execresult
insert into user_info (name, username, password, extra_data, email) VALUES (?, ?, ?, ?, ?)
`

type CreateNewUserInfoParams struct {
	Name      string         `json:"name"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	ExtraData sql.NullString `json:"extra_data"`
	Email     string         `json:"email"`
}

func (q *Queries) CreateNewUserInfo(ctx context.Context, arg CreateNewUserInfoParams) (sql.Result, error) {
	return q.exec(ctx, q.createNewUserInfoStmt, createNewUserInfo,
		arg.Name,
		arg.Username,
		arg.Password,
		arg.ExtraData,
		arg.Email,
	)
}

const getUserInfoById = `-- name: GetUserInfoById :one
select id, name, username, password, extra_data, created_at, updated_at, email from user_info where id = ? limit 1
`

func (q *Queries) GetUserInfoById(ctx context.Context, id int64) (UserInfo, error) {
	row := q.queryRow(ctx, q.getUserInfoByIdStmt, getUserInfoById, id)
	var i UserInfo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Password,
		&i.ExtraData,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
	)
	return i, err
}

const updateUserInfoByUsername = `-- name: UpdateUserInfoByUsername :execresult
update user_info set extra_data = ? where username = ?
`

type UpdateUserInfoByUsernameParams struct {
	ExtraData sql.NullString `json:"extra_data"`
	Username  string         `json:"username"`
}

func (q *Queries) UpdateUserInfoByUsername(ctx context.Context, arg UpdateUserInfoByUsernameParams) (sql.Result, error) {
	return q.exec(ctx, q.updateUserInfoByUsernameStmt, updateUserInfoByUsername, arg.ExtraData, arg.Username)
}
