package repositories

import (
	"context"
	"github.com/arioprima/jobseekers_api/models"
	"github.com/arioprima/jobseekers_api/schemas"
	"gorm.io/gorm"
	"net/http"
)

type ResendOtpRepository interface {
	ResendOtp(ctx context.Context, tx *gorm.DB, req *schemas.SchemaDataUser) (string, *schemas.SchemaDatabaseError)
}

type ResendOtpRepositoryImpl struct {
	DB *gorm.DB
}

func NewResendOtpRepositoryImpl(db *gorm.DB) ResendOtpRepository {
	return &ResendOtpRepositoryImpl{DB: db}
}

func (r *ResendOtpRepositoryImpl) ResendOtp(ctx context.Context, tx *gorm.DB, req *schemas.SchemaDataUser) (string, *schemas.SchemaDatabaseError) {
	var (
		otpCode models.OtpCode
		err     error
	)
	if tx == nil {
		tx = r.DB.WithContext(ctx).Debug().Begin()
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Where("user_id = ?", req.ID).First(&otpCode).Error
	if err != nil {
		return "", &schemas.SchemaDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_01",
		}
	}

	//update otp code
	err = tx.Where("user_id = ?", req.ID).Updates(models.OtpCode{
		Code: req.OtpCode,
	}).Error

	if err != nil {
		return "", &schemas.SchemaDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_02",
		}
	}

	return otpCode.Code, nil
}
