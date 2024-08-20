package models

import "time"

type User struct {
	ID          int       `json:"id" db:"id"`
	INN         string    `json:"inn" db:"inn"`
	Username    string    `json:"username" db:"username"`
	Password    string    `json:"password" db:"password"`
	CompanyName string    `json:"company_name" db:"company_name"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	IsAdmin     bool      `json:"is_admin" db:"is_admin"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type UserCreateRequest struct {
	INN         string `json:"inn"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	CompanyName string `json:"company_name"`
	IsActive    bool   `json:"is_active"`
	IsAdmin     bool   `json:"is_admin"`
}

type UserUpdateRequest struct {
	INN         *string `json:"inn,omitempty"`
	Username    *string `json:"username,omitempty"`
	Password    *string `json:"password,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
	IsAdmin     *bool   `json:"is_admin,omitempty"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
