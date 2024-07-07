package models

import "time"

type User struct {
	ID           string    `json:"id" validate:"uuid"`
	BiodataId    string    `json:"biodata_id" validate:"required,uuid" gore:"column:biodata_id"`
	Password     string    `json:"password" validate:"required"`
	IsActive     bool      `json:"is_active,omitempty" gorm:"default:true"`
	IsVerified   bool      `json:"is_verified,omitempty"  gorm:"default:false"`
	ProfileImage string    `json:"profile_image,omitempty" gorm:"column:profile_image"`
	RoleId       string    `json:"role_id" validate:"required,uuid" gorm:"column:role_id"`
	Summary      string    `json:"summary,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
