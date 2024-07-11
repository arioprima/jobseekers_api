package services

import (
	"context"
	"github.com/arioprima/jobseekers_api/config"
	"github.com/arioprima/jobseekers_api/models"
	"github.com/arioprima/jobseekers_api/pkg"
	repositories "github.com/arioprima/jobseekers_api/repositories/auth"
	"github.com/arioprima/jobseekers_api/schemas"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"time"
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
		log.Println("kirimErorkeHandler", err)
		return nil, err
	}

	configs, _ := config.LoadConfig(".")
	accessTokenData := map[string]interface{}{
		"id":        res.ID,
		"email":     res.Biodata.Email,
		"firstname": res.Biodata.Firstname,
		"lastname":  res.Biodata.Lastname,
		"role_id":   res.Role.ID,
		"role_name": res.Role.Name,
	}

	accessToken, tokenErr := pkg.GenerateToken(accessTokenData, configs.TokenSecret, configs.TokenExpired)

	if tokenErr != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: 500,
			Type: "error_04",
		}
	}

	res.Auth.AccessToken = accessToken
	res.Auth.Type = "Bearer"
	res.Auth.ExpiredAt = pkg.CalculateExpiration(time.Now().Add(configs.TokenExpired).Unix())
	return res, err
}
