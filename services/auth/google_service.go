package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/arioprima/jobseekers_api/config"
	repositories "github.com/arioprima/jobseekers_api/repositories/auth"
	"github.com/arioprima/jobseekers_api/schemas"
	"gorm.io/gorm"
	"time"
)

type ServiceGoogle interface {
	LoginGoogleService(ctx context.Context, tx *gorm.DB, data *schemas.SchemaDataUser) (*schemas.LoginUserResponse, *schemas.SchemaDatabaseError)
}

type serviceGoogleImpl struct {
	repository repositories.RepositoryGoogle
}

func NewServiceGoogleImpl(repository repositories.RepositoryGoogle) ServiceGoogle {
	return &serviceGoogleImpl{
		repository: repository,
	}
}

func (s *serviceGoogleImpl) LoginGoogleService(ctx context.Context, tx *gorm.DB, data *schemas.SchemaDataUser) (*schemas.LoginUserResponse, *schemas.SchemaDatabaseError) {
	configs, _ := config.LoadConfig(".")

	hashed := sha256.New()
	hashed.Write([]byte(configs.TokenSecret + time.Now().String()))
	token := hex.EncodeToString(hashed.Sum(nil))
	data.Token = token

	_, err := s.repository.LoginGoogle(ctx, tx, data)
	res := schemas.LoginUserResponse{
		ID:           data.ID,
		Firstname:    data.Firstname,
		Lastname:     data.Lastname,
		Email:        data.Email,
		RoleId:       data.RoleId,
		RoleName:     "user",
		ProfileImage: &data.ProfileImage,
	}

	return &res, err
}
