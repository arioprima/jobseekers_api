package models

import "time"

type UserSession struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	UserID    string     `gorm:"not null"`
	Token     string     `gorm:"not null"`
	LastLogin time.Time  `gorm:"default:null"`
	ExpiredAt time.Time  `gorm:"not null"`
	CreatedAt time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"default:current_timestamp on update current_timestamp"`
	DeletedAt *time.Time `gorm:"index"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}
