// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: face.sql

package store

import (
	"context"
	"database/sql"
)

const createFaceByUserId = `-- name: CreateFaceByUserId :execresult
insert into user_face (user_id, data)
VALUES (?, ?)
`

type CreateFaceByUserIdParams struct {
	UserID int64          `json:"user_id"`
	Data   sql.NullString `json:"data"`
}

func (q *Queries) CreateFaceByUserId(ctx context.Context, arg CreateFaceByUserIdParams) (sql.Result, error) {
	return q.exec(ctx, q.createFaceByUserIdStmt, createFaceByUserId, arg.UserID, arg.Data)
}

const getFaceByUserId = `-- name: GetFaceByUserId :one
select user_id, data, created_at, updated_at
from user_face
where ` + "`" + `user_id` + "`" + ` = ?
`

func (q *Queries) GetFaceByUserId(ctx context.Context, userID int64) (UserFace, error) {
	row := q.queryRow(ctx, q.getFaceByUserIdStmt, getFaceByUserId, userID)
	var i UserFace
	err := row.Scan(
		&i.UserID,
		&i.Data,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFaceByUserIdForUpdate = `-- name: GetFaceByUserIdForUpdate :one
select user_id, data, created_at, updated_at
from user_face
where user_id = ? for
update
`

func (q *Queries) GetFaceByUserIdForUpdate(ctx context.Context, userID int64) (UserFace, error) {
	row := q.queryRow(ctx, q.getFaceByUserIdForUpdateStmt, getFaceByUserIdForUpdate, userID)
	var i UserFace
	err := row.Scan(
		&i.UserID,
		&i.Data,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateFaceByUserId = `-- name: UpdateFaceByUserId :execresult
update user_face
set data = ?
where user_id = ?
`

type UpdateFaceByUserIdParams struct {
	Data   sql.NullString `json:"data"`
	UserID int64          `json:"user_id"`
}

func (q *Queries) UpdateFaceByUserId(ctx context.Context, arg UpdateFaceByUserIdParams) (sql.Result, error) {
	return q.exec(ctx, q.updateFaceByUserIdStmt, updateFaceByUserId, arg.Data, arg.UserID)
}