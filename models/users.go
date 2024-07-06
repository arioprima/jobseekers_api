package models

import "time"

type User struct {
	ID         string    `json:"id" validate:"uuid"`
	BiodataId  string    `json:"biodata_id" validate:"required,uuid" gore:"column:biodata_id"`
	Password   string    `json:"password" validate:"required"`
	IsActive   bool      `json:"is_active,omitempty" gorm:"default:true,column:is_active"`
	IsVerified bool      `json:"is_verified,omitempty" gorm:"default:false,column:is_verified"`
	RoleId     string    `json:"role_id" validate:"required,uuid" gorm:"column:role_id"`
	Summary    string    `json:"summary,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
