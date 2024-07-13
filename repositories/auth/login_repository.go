package repositories

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/arioprima/jobseekers_api/config"
	"github.com/arioprima/jobseekers_api/models"
	"github.com/arioprima/jobseekers_api/pkg"
	"github.com/arioprima/jobseekers_api/schemas"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type RepositoryLogin interface {
	Login(ctx context.Context, tx *gorm.DB, req *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError)
}

type repositoryLoginImpl struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func NewRepositoryLoginImpl(log *logrus.Logger, db *gorm.DB) RepositoryLogin {
	return &repositoryLoginImpl{
		DB:  db,
		Log: log,
	}
}

func (r *repositoryLoginImpl) Login(ctx context.Context, tx *gorm.DB, req *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError) {
	//TODO implement me
	var (
		user models.ModelAuth
		err  error
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

	if err = tx.Joins("LEFT JOIN biodata ON users.biodata_id = biodata.id").
		Joins("LEFT JOIN user_roles ON users.role_id = user_roles.id").
		Preload("Biodata").
		Preload("Role").
		Where("biodata.email = ?", req.Email).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &schemas.SchemaDatabaseError{
				Code: http.StatusNotFound,
				Type: "error_01",
			}
		}
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_02",
		}
	}

	comparePassword := pkg.ComparePassword(user.Password, req.Password)
	if comparePassword != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusUnauthorized,
			Type: "error_03",
		}
	}

	// Create user session
	configs, _ := config.LoadConfig(".")
	hashed := sha256.New()
	hashed.Write([]byte(configs.TokenSecret + time.Now().String()))
	token := hex.EncodeToString(hashed.Sum(nil))
	user.Token = token

	session := models.UserSession{
		UserID:    user.ID,
		Token:     token,
		LastLogin: time.Now(),
		ExpiredAt: pkg.CalculateExpiration(time.Now().Add(configs.TokenExpired).Unix()),
	}

	if err := tx.Create(&session).Error; err != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_02",
		}
	}

	return &user, nil
}
