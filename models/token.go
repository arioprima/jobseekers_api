package models

import "time"

type TokenAuth struct {
	AccessToken string    `json:"token"`
	Type        string    `json:"type"`
	ExpiredAt   time.Time `json:"expired_at"`
}
