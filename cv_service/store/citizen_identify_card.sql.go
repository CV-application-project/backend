// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: citizen_identify_card.sql

package store

import (
	"context"
	"database/sql"
)

const createCICByUserId = `-- name: CreateCICByUserId :execresult
insert into citizen_identify_card (user_id, data)
VALUES (?, ?)
`

type CreateCICByUserIdParams struct {
	UserID int64          `json:"user_id"`
	Data   sql.NullString `json:"data"`
}

func (q *Queries) CreateCICByUserId(ctx context.Context, arg CreateCICByUserIdParams) (sql.Result, error) {
	return q.exec(ctx, q.createCICByUserIdStmt, createCICByUserId, arg.UserID, arg.Data)
}

const getCICByUserId = `-- name: GetCICByUserId :one
select user_id, data, created_at, updated_at
from citizen_identify_card
where ` + "`" + `user_id` + "`" + ` = ?
`

func (q *Queries) GetCICByUserId(ctx context.Context, userID int64) (CitizenIdentifyCard, error) {
	row := q.queryRow(ctx, q.getCICByUserIdStmt, getCICByUserId, userID)
	var i CitizenIdentifyCard
	err := row.Scan(
		&i.UserID,
		&i.Data,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCICByUserIdForUpdate = `-- name: GetCICByUserIdForUpdate :one
select user_id, data, created_at, updated_at
from citizen_identify_card
where user_id = ? for
update
`

func (q *Queries) GetCICByUserIdForUpdate(ctx context.Context, userID int64) (CitizenIdentifyCard, error) {
	row := q.queryRow(ctx, q.getCICByUserIdForUpdateStmt, getCICByUserIdForUpdate, userID)
	var i CitizenIdentifyCard
	err := row.Scan(
		&i.UserID,
		&i.Data,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCICByUserId = `-- name: UpdateCICByUserId :execresult
update citizen_identify_card
set data = ?
where user_id = ?
`

type UpdateCICByUserIdParams struct {
	Data   sql.NullString `json:"data"`
	UserID int64          `json:"user_id"`
}

func (q *Queries) UpdateCICByUserId(ctx context.Context, arg UpdateCICByUserIdParams) (sql.Result, error) {
	return q.exec(ctx, q.updateCICByUserIdStmt, updateCICByUserId, arg.Data, arg.UserID)
}
