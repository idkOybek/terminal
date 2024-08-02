package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	INN       string    `json:"inn" db:"inn"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	IsAdmin   bool      `json:"is_admin" db:"is_admin"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserCreateRequest struct {
	INN      string `json:"inn"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserUpdateRequest struct {
	INN      string `json:"inn"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}