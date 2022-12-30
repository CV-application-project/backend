// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package store

import (
	"database/sql"
	"time"
)

type User struct {
	ID         int64          `json:"id"`
	Name       string         `json:"name"`
	EmployeeID string         `json:"employee_id"`
	Password   string         `json:"password"`
	Phone      sql.NullString `json:"phone"`
	Address    sql.NullString `json:"address"`
	Gender     sql.NullString `json:"gender"`
	Department sql.NullString `json:"department"`
	Position   sql.NullString `json:"position"`
	Role       sql.NullString `json:"role"`
	Data       sql.NullString `json:"data"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	Email      string         `json:"email"`
	FrontCard  sql.NullString `json:"front_card"`
	BackCard   sql.NullString `json:"back_card"`
}

type UserToken struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
