package services

import (
	"context"
	"github.com/arioprima/jobseekers_api/models"
	repositories "github.com/arioprima/jobseekers_api/repositories/auth"
	"github.com/arioprima/jobseekers_api/schemas"
	"gorm.io/gorm"
)

type ServiceVerifyEmail interface {
	VerifyEmailService(ctx context.Context, tx *gorm.DB, userID string, otp string) (*models.OtpCode, *schemas.SchemaDatabaseError)
}

type serviceVerifyEmailImpl struct {
	repository repositories.VerifyEmailRepository
}

func NewServiceVerifyEmailImpl(repository repositories.VerifyEmailRepository) ServiceVerifyEmail {
	return &serviceVerifyEmailImpl{
		repository: repository,
	}
}

func (s *serviceVerifyEmailImpl) VerifyEmailService(ctx context.Context, tx *gorm.DB, userID string, otp string) (*models.OtpCode, *schemas.SchemaDatabaseError) {
	//TODO implement me
	res, err := s.repository.VerifyEmail(ctx, tx, userID, otp)
	if err != nil {
		return nil, err
	}
	return res, err
}
