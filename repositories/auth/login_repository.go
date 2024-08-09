package repositories

import (
	"context"
	"errors"
	"github.com/arioprima/jobseekers_api/models"
	"github.com/arioprima/jobseekers_api/pkg"
	"github.com/arioprima/jobseekers_api/schemas"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type RepositoryLogin interface {
	Login(ctx context.Context, tx *gorm.DB, req *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError)
}

type repositoryLoginImpl struct {
	DB *gorm.DB
}

func NewRepositoryLoginImpl(db *gorm.DB) RepositoryLogin {
	return &repositoryLoginImpl{
		DB: db,
	}
}

func (r *repositoryLoginImpl) Login(ctx context.Context, tx *gorm.DB, req *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError) {
	//TODO implement me
	var (
		user models.ModelAuth
		bio  models.Biodata
		err  error
	)

	if tx == nil {
		tx = r.DB.WithContext(ctx).Begin()
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
	if checkUserAccount.RowsAffected < 1 {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_01",
		}
	}

	if err = tx.Joins("LEFT JOIN biodata b ON users.biodata_id = b.id").
		Joins("LEFT JOIN user_roles ur ON users.role_id = ur.id").
		Preload("Biodata").
		Preload("Role").
		Where("b.email = ?", req.Email).
		Where("users.is_active = ?", true).
		Where("users.is_verified = ?", true).
		Where("users.deleted_at IS NULL").
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &schemas.SchemaDatabaseError{
				Code: http.StatusUnauthorized,
				Type: "error",
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

	session := models.UserSession{
		UserID:    user.ID,
		Token:     req.Token,
		LastLogin: time.Now(),
		ExpiredAt: req.ExpiredAt,
	}

	if err := tx.Create(&session).Error; err != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_02",
		}
	}

	return &user, nil
}
