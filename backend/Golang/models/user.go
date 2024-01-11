package models

import "time"

type User struct {
	UserID            string    `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	FirstUser         bool      `json:"first_user"`
	IsActive          bool      `json:"is_active"`
	IsVerified        bool      `json:"is_verified"`
	VerificationToken string    `json:"verification_token"`
	RoleID            string    `json:"role_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type AdminUser struct {
	AdminID    string `json:"id"`
	BirthPlace string `json:"birth_place"`
	BirthDate  string `json:"date_of_birth"`
	UserID     string `json:"user_id"`
}

type RegisterInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleID    string `json:"role_id"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyInput struct {
	Email     string `json:"email"`
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}

type LoginResponse struct {
	Email     string `json:"email"`
	FirstUser bool   `json:"first_user"`
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}

type UserResponse struct {
	UserID    string `json:"id"`
	Email     string `json:"email"`
	RoleID    string `json:"role_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
