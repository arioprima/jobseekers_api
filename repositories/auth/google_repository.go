package repositories

import (
	"context"
	"github.com/arioprima/jobseekers_api/models"
	"github.com/arioprima/jobseekers_api/pkg"
	"github.com/arioprima/jobseekers_api/schemas"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type RepositoryGoogle interface {
	LoginGoogle(ctx context.Context, tx *gorm.DB, req *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError)
}

type repositoryGoogleImpl struct {
	DB *gorm.DB
}

func NewRepositoryGoogleImpl(db *gorm.DB) RepositoryGoogle {
	return &repositoryGoogleImpl{
		DB: db,
	}
}

func (r *repositoryGoogleImpl) LoginGoogle(ctx context.Context, tx *gorm.DB, req *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError) {
	var (
		user models.ModelAuth
		bio  models.Biodata
		err  error
	)

	session := models.UserSession{
		UserID:    req.ID,
		Token:     req.Token,
		LastLogin: time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}

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

	checkUserAccount := tx.Where("email = ?", req.Email).First(&bio)
	if checkUserAccount.Error != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_01",
		}
	}

	if checkUserAccount.RowsAffected > 0 {
		if err := tx.Create(&session).Error; err != nil {
			return nil, &schemas.SchemaDatabaseError{
				Code: http.StatusInternalServerError,
				Type: "error_02",
			}
		}
		return &user, nil
	}

	BioId := pkg.GenerateUUID()

	bio.ID = BioId
	bio.Firstname = req.Firstname
	bio.Lastname = req.Lastname
	bio.Email = req.Email

	if err = tx.Create(&bio).Error; err != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_03",
		}
	}

	user.ID = req.ID
	user.BiodataId = BioId
	user.IsActive = true
	user.IsVerified = true
	user.ProfileImage = &req.ProfileImage
	user.RoleId = req.RoleId

	if err = tx.Create(&user).Error; err != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_04",
		}
	}

	if err := tx.Create(&session).Error; err != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_02",
		}
	}

	return &user, nil
}
