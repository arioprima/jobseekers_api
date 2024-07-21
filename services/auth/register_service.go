package services

import (
	"context"
	"github.com/arioprima/jobseekers_api/pkg"
	"github.com/sirupsen/logrus"
	"log"

	repositories "github.com/arioprima/jobseekers_api/repositories/auth"
	"github.com/arioprima/jobseekers_api/schemas"
	"gorm.io/gorm"
)

type ServiceRegister interface {
	RegisterService(ctx context.Context, tx *gorm.DB, input *schemas.SchemaDataUser) (*schemas.OtpEmailResponse, *schemas.SchemaDatabaseError)
}

type serviceRegisterImpl struct {
	repository repositories.RegisterRepository
}

func NewServiceRegisterImpl(repository repositories.RegisterRepository) ServiceRegister {
	return &serviceRegisterImpl{
		repository: repository,
	}
}

func (s *serviceRegisterImpl) RegisterService(ctx context.Context, tx *gorm.DB, input *schemas.SchemaDataUser) (*schemas.OtpEmailResponse, *schemas.SchemaDatabaseError) {
	var schema schemas.SchemaDataUser

	switch input.RoleId {
	case "019047ca-f542-7182-8b6b-7978f905dfe7":
		input.RoleId = "019047ca-f542-7182-8b6b-7978f905dfe7"
	case "019047ca-f542-71fe-9de6-c4919ed5c9ff":
		input.RoleId = "019047ca-f542-71fe-9de6-c4919ed5c9ff"
	default:
		input.RoleId = "01908d0f-289d-7fd7-9143-d9525f8bc74d"
	}

	log.Println("Role ID: ", input.RoleId)

	userId := pkg.GenerateUUID()
	bioId := pkg.GenerateUUID()
	// Generate OTP Code
	otpCode := pkg.GenerateOtp()

	schema.ID = userId
	schema.BiodataId = bioId
	schema.Firstname = input.Firstname
	schema.Lastname = input.Lastname
	schema.Email = input.Email
	schema.RoleId = input.RoleId
	schema.OtpCode = otpCode

	hashedPassword := pkg.HashPassword(input.Password)
	schema.Password = hashedPassword

	_, err := s.repository.Register(ctx, tx, &schema)
	if err != nil {
		return nil, err
	}

	// Start sending email in a goroutine
	go func() {
		logrus.Info("Send Email Start")
		name := input.Firstname + " " + input.Lastname
		pkg.SendEmail(name, input.Email, otpCode, "template_register")
		logrus.Info("Send Email Done")
	}()

	res := &schemas.OtpEmailResponse{
		ID:      userId,
		Email:   input.Email,
		OtpCode: otpCode,
	}

	return res, nil
}
