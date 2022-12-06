// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package store

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createNewUserInfoStmt, err = db.PrepareContext(ctx, createNewUserInfo); err != nil {
		return nil, fmt.Errorf("error preparing query CreateNewUserInfo: %w", err)
	}
	if q.createNewUserTokenByUserIdStmt, err = db.PrepareContext(ctx, createNewUserTokenByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query CreateNewUserTokenByUserId: %w", err)
	}
	if q.getUserByUsernameOrEmailStmt, err = db.PrepareContext(ctx, getUserByUsernameOrEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByUsernameOrEmail: %w", err)
	}
	if q.getUserInfoByIdStmt, err = db.PrepareContext(ctx, getUserInfoById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserInfoById: %w", err)
	}
	if q.getUserTokenByUserIdStmt, err = db.PrepareContext(ctx, getUserTokenByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserTokenByUserId: %w", err)
	}
	if q.updateUserInfoByIdStmt, err = db.PrepareContext(ctx, updateUserInfoById); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserInfoById: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createNewUserInfoStmt != nil {
		if cerr := q.createNewUserInfoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createNewUserInfoStmt: %w", cerr)
		}
	}
	if q.createNewUserTokenByUserIdStmt != nil {
		if cerr := q.createNewUserTokenByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createNewUserTokenByUserIdStmt: %w", cerr)
		}
	}
	if q.getUserByUsernameOrEmailStmt != nil {
		if cerr := q.getUserByUsernameOrEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByUsernameOrEmailStmt: %w", cerr)
		}
	}
	if q.getUserInfoByIdStmt != nil {
		if cerr := q.getUserInfoByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserInfoByIdStmt: %w", cerr)
		}
	}
	if q.getUserTokenByUserIdStmt != nil {
		if cerr := q.getUserTokenByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserTokenByUserIdStmt: %w", cerr)
		}
	}
	if q.updateUserInfoByIdStmt != nil {
		if cerr := q.updateUserInfoByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserInfoByIdStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                             DBTX
	tx                             *sql.Tx
	createNewUserInfoStmt          *sql.Stmt
	createNewUserTokenByUserIdStmt *sql.Stmt
	getUserByUsernameOrEmailStmt   *sql.Stmt
	getUserInfoByIdStmt            *sql.Stmt
	getUserTokenByUserIdStmt       *sql.Stmt
	updateUserInfoByIdStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                             tx,
		tx:                             tx,
		createNewUserInfoStmt:          q.createNewUserInfoStmt,
		createNewUserTokenByUserIdStmt: q.createNewUserTokenByUserIdStmt,
		getUserByUsernameOrEmailStmt:   q.getUserByUsernameOrEmailStmt,
		getUserInfoByIdStmt:            q.getUserInfoByIdStmt,
		getUserTokenByUserIdStmt:       q.getUserTokenByUserIdStmt,
		updateUserInfoByIdStmt:         q.updateUserInfoByIdStmt,
	}
}