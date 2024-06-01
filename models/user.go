package models

import (
	"database/sql"
	"time"
)

// create UserInfo struct
type UserInfo struct {
	Email     string `json:"email"`
	FullName  string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Picture   string `json:"picture"`
}

// create User struct

type User struct {
	ID       int `json:"id"`
	UserInfo UserInfo
	Rank     int `json:"rank"`
}

type Token struct {
	ID        int          `db:"id"`
	UserID    int          `db:"user_id"`
	Name      string       `db:"name"`
	Token     string       `db:"token"`
	CreatedAt time.Time    `db:"created_at"`
	RevokedAt sql.NullTime `db:"revoked_at,omitempty"`
}
