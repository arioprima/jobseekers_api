package models

import "time"

type UserRole struct {
	ID        string    `json:"id" validate:"uuid"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *UserRole) tableName() string {
	return "user_roles"
}
