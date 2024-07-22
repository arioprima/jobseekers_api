package repositories

import (
	"context"
	"github.com/arioprima/jobseekers_api/models"
	"github.com/arioprima/jobseekers_api/schemas"
	"gorm.io/gorm"
	"time"
)

type VerifyEmailRepository interface {
	VerifyEmail(ctx context.Context, tx *gorm.DB, userID string, otp string) (*models.OtpCode, *schemas.SchemaDatabaseError)
}

type verifyEmailRepositoryImpl struct {
	DB *gorm.DB
}

func NewVerifyEmailRepositoryImpl(db *gorm.DB) VerifyEmailRepository {
	return &verifyEmailRepositoryImpl{
		DB: db,
	}
}

func (r *verifyEmailRepositoryImpl) VerifyEmail(ctx context.Context, tx *gorm.DB, userID string, otp string) (*models.OtpCode, *schemas.SchemaDatabaseError) {
	var (
		user    models.User
		otpCode models.OtpCode
		err     error
	)

	if tx == nil {
		tx = r.DB.WithContext(ctx).Debug().Begin()
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Where("user_id = ? AND code = ? AND expired_at >", userID, otp, time.Now()).First(&otpCode).Error
	if err != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: 400,
			Type: "error_01",
		}
	}
	//update is_verified di table user
	user.IsVerified = true
	_ = tx.Model(&user).Where("id = ?", userID).Update("is_verified", true).Error

	return &otpCode, nil
}
