package services

import (
	"context"
	"github.com/arioprima/jobseekers_api/models"
	repositories "github.com/arioprima/jobseekers_api/repositories/auth"
	"github.com/arioprima/jobseekers_api/schemas"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ServiceLogin interface {
	LoginService(ctx context.Context, tx *gorm.DB, input *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError)
}

type serviceLoginImpl struct {
	repository repositories.RepositoryLogin
	Log        *logrus.Logger
}

func NewServiceLoginImpl(repository repositories.RepositoryLogin, log *logrus.Logger) ServiceLogin {
	return &serviceLoginImpl{
		repository: repository,
		Log:        log,
	}
}

func (s *serviceLoginImpl) LoginService(ctx context.Context, tx *gorm.DB, input *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError) {
	//TODO implement me
	var schema schemas.SchemaDataUser
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.repository.Login(ctx, tx, &schema)
	if err != nil {
		return nil, err
	}
	return res, err
}
