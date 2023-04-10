package model

import (
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	ID        int            `json:"user_id"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Balance   int            `json:"balance"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserTokenResponse struct {
	Token string `json:"token"`
}

type UserSignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserEditRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
