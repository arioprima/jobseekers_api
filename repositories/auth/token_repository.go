package repositories

import (
	"github.com/arioprima/jobseekers_api/models"
	"gorm.io/gorm"
)

func FinByToken(userId string, db *gorm.DB) (string, error) {
	var userToken []models.UserSession
	err := db.Debug().Order("created_at desc").
		Select("token").Where("user_id = ? and expired_at >= NOW()", userId).
		First(&userToken).Error

	if err != nil {
		return "", err
	}

	return userToken[0].Token, nil
}
