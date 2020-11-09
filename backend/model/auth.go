package model

import "time"

// Login is used to login & get jwt
type Login struct {
	UserID   string `json:"userID" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginSuccessMessage is returned when login successed
type LoginSuccessMessage struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

// UnauthorizedMessage is returned when login failed
type UnauthorizedMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// User to login
type User struct {
	UserID string
}
