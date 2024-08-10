package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/arioprima/jobseekers_api/config"
	"github.com/arioprima/jobseekers_api/models"
	"github.com/arioprima/jobseekers_api/pkg"
	repositories "github.com/arioprima/jobseekers_api/repositories/auth"
	"github.com/arioprima/jobseekers_api/schemas"
	"gorm.io/gorm"
	"time"
)

type ServiceLogin interface {
	LoginService(ctx context.Context, tx *gorm.DB, input *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError, *models.TokenAuth)
}

type serviceLoginImpl struct {
	repository repositories.RepositoryLogin
}

func NewServiceLoginImpl(repository repositories.RepositoryLogin) ServiceLogin {
	return &serviceLoginImpl{
		repository: repository,
	}
}

func (s *serviceLoginImpl) LoginService(ctx context.Context, tx *gorm.DB, input *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError, *models.TokenAuth) {
	//TODO implement me
	configs, _ := config.LoadConfig(".")

	hashed := sha256.New()
	hashed.Write([]byte(configs.TokenSecret + time.Now().String()))
	token := hex.EncodeToString(hashed.Sum(nil))
	expiredAt := pkg.CalculateExpiration(time.Now().Add(configs.TokenExpired).Unix())

	var schema schemas.SchemaDataUser
	schema.Email = input.Email
	schema.Password = input.Password
	schema.Token = token
	schema.ExpiredAt = expiredAt

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
		"token":         token,
	}

	accessToken, _ := pkg.GenerateToken(accessTokenData, configs.TokenSecret, configs.TokenExpired)

	authToken := models.TokenAuth{
		AccessToken: accessToken,
		Type:        "Bearer",
		ExpiredAt:   expiredAt,
	}

	if res.ProfileImage != nil && *res.ProfileImage == "" {
		res.ProfileImage = nil
	}

	return res, err, &authToken
}
