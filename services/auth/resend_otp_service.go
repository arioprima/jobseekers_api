package services

import (
	"context"
	"github.com/arioprima/jobseekers_api/pkg"
	repositories "github.com/arioprima/jobseekers_api/repositories/auth"
	"github.com/arioprima/jobseekers_api/schemas"
)

type ServiceResendOtp interface {
	ResendOtp(ctx context.Context, input *schemas.SchemaDataUser, userId string) (string, *schemas.SchemaDatabaseError)
}

type serviceResendOtpImpl struct {
	repository repositories.ResendOtpRepository
}

func NewServiceResendOtpImpl(repository repositories.ResendOtpRepository) ServiceResendOtp {
	return &serviceResendOtpImpl{
		repository: repository,
	}
}

func (s *serviceResendOtpImpl) ResendOtp(ctx context.Context, input *schemas.SchemaDataUser, userId string) (string, *schemas.SchemaDatabaseError) {
	var schema schemas.SchemaDataUser

	otpCode := pkg.GenerateOtp()
	go func() {
		name := input.Firstname + " " + input.Lastname
		pkg.SendEmail(name, input.Email, otpCode, "template_resend")
	}()

	schema.ID = userId
	schema.OtpCode = otpCode

	_, err := s.repository.ResendOtp(ctx, nil, &schema)
	if err != nil {
		return "", err
	}

	return otpCode, nil
}
