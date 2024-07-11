package models

import "time"

type Auth struct {
	AccessToken string    `json:"token"`
	Type        string    `json:"type"`
	ExpiredAt   time.Time `json:"expired_at"`
}

type ModelAuth struct {
	ID           string    `json:"id" gorm:"primaryKey;column:id"`
	BiodataId    string    `json:"biodata_id,omitempty" gorm:"column:biodata_id"`
	Biodata      Biodata   `json:"biodata,omitempty" gorm:"foreignKey:BiodataId;references:ID"`
	Password     string    `json:"password,omitempty" gorm:"column:password;type:varchar(255)"`
	IsActive     bool      `json:"is_active,omitempty" gorm:"column:is_active;default:true"`
	IsVerified   bool      `json:"is_verified,omitempty" gorm:"column:is_verified;default:false"`
	ProfileImage string    `json:"profile_image,omitempty" gorm:"column:profile_image"`
	RoleId       string    `json:"role_id,omitempty" gorm:"column:role_id"`
	Role         UserRole  `json:"role,omitempty" gorm:"foreignKey:RoleId;references:ID"`
	Summary      string    `json:"summary,omitempty" gorm:"column:summary"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
	Auth         Auth      `json:"auth,omitempty" gorm:"-"`
	Token        string    `json:"token,omitempty" gorm:"-"`
}

func (auth *ModelAuth) TableName() string {
	return "users"
}
