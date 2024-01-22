package models

import "time"

type User struct {
	UserID            string    `json:"id"`
	FirstName         string    `json:"first_name" validate:"required"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email" validate:"required,email"`
	Password          string    `json:"password" validate:"required,min=4,max=32"`
	FirstUser         bool      `json:"first_user"`
	IsActive          bool      `json:"is_active"`
	IsVerified        bool      `json:"is_verified"`
	VerificationToken string    `json:"verification_token"`
	RoleID            string    `json:"role_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type AdminUser struct {
	AdminID    string    `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	BirthPlace string    `json:"birth_place"`
	BirthDate  string    `json:"date_of_birth"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	UserID     string    `json:"user_id" validate:"required"`
	RoleID     string    `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RecruiterUser struct {
	RecruiterID       string `json:"id"`
	Phone             string `json:"phone"`
	Address           string `json:"address"`
	CompanyName       string `json:"company_name"`
	CompanyLogo       string `json:"company_logo"`
	CompanyDesc       string `json:"company_desc"`
	CompanyCategoryID string `json:"company_category_id"`
}

type JobSeekerUser struct {
	JobSeekerID string `json:"id"`
	BirthPlace  string `json:"birth_place"`
	BirthDate   string `json:"date_of_birth"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
}

type AdminInput struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	BirthPlace string `json:"birth_place"`
	BirthDate  string `json:"date_of_birth"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	FirstUser  bool   `json:"first_user"`
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
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserResponse struct {
	UserID    string `json:"id"`
	Email     string `json:"email"`
	RoleID    string `json:"role_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FirstUser bool   `json:"first_user"`
}
