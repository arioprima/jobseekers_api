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
	"time"
)

type ServiceLogin interface {
	LoginService(ctx context.Context, tx *gorm.DB, input *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError, *models.TokenAuth)
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

func (s *serviceLoginImpl) LoginService(ctx context.Context, tx *gorm.DB, input *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError, *models.TokenAuth) {
	//TODO implement me
	configs, _ := config.LoadConfig(".")
	var schema schemas.SchemaDataUser
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.repository.Login(ctx, tx, &schema)
	if err != nil {
		return nil, err, nil
	}

	accessTokenData := map[string]interface{}{
		"id":            res.ID,
		"email":         res.Biodata.Email,
		"firstname":     res.Biodata.Firstname,
		"lastname":      res.Biodata.Lastname,
		"role_id":       res.Role.ID,
		"role_name":     res.Role.Name,
		"profile_image": res.ProfileImage,
		"token":         res.Token,
	}

	accessToken, _ := pkg.GenerateToken(accessTokenData, configs.TokenSecret, configs.TokenExpired)

	authToken := models.TokenAuth{
		AccessToken: accessToken,
		Type:        "Bearer",
		ExpiredAt:   pkg.CalculateExpiration(time.Now().Add(configs.TokenExpired).Unix()),
	}

	if res.ProfileImage != nil && *res.ProfileImage == "" {
		res.ProfileImage = nil
	}

	return res, err, &authToken
}
