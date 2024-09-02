package models

import "time"

type User struct {
	UserID      string    `json:"user_id"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_no"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateUserReq struct {
	FirstName   string `json:"firstname" validate:"required,alpha"`
	LastName    string `json:"lastname" validate:"required,alpha"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_no" validate:"required"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponseInfo struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
