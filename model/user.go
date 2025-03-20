package model

import "github.com/google/uuid"

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserParam struct {
	UserID   uuid.UUID `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
	RoleID   int       `json:"-"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UserID uuid.UUID `json:"userID"`
	Token  string    `json:"token"`
	RoleID int       `json:"role_id"`
}

type GoogleLogin struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}
