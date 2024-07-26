package repositories

import (
	"context"
	"github.com/arioprima/jobseekers_api/models"
	"github.com/arioprima/jobseekers_api/pkg"
	"github.com/arioprima/jobseekers_api/schemas"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

type RegisterRepository interface {
	Register(ctx context.Context, tx *gorm.DB, req *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError)
}

type registerRepositoryImpl struct {
	DB *gorm.DB
}

func NewRegisterRepositoryImpl(db *gorm.DB) RegisterRepository {
	return &registerRepositoryImpl{
		DB: db,
	}
}

func (r *registerRepositoryImpl) Register(ctx context.Context, tx *gorm.DB, req *schemas.SchemaDataUser) (*models.ModelAuth, *schemas.SchemaDatabaseError) {
	var (
		user models.ModelAuth
		bio  models.Biodata
		err  error
		wg   sync.WaitGroup
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

	errChan := make(chan error, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		checkUserAccount := tx.Where("email = ?", req.Email).First(&bio)
		if checkUserAccount.RowsAffected > 0 {
			errChan <- &schemas.SchemaDatabaseError{
				Code: http.StatusConflict,
				Type: "error_01",
			}
		}
	}()

	wg.Wait()
	close(errChan)

	if err = <-errChan; err != nil {
		return nil, err.(*schemas.SchemaDatabaseError)
	}

	user.ID = req.ID
	user.BiodataId = req.BiodataId
	user.Biodata.ID = req.BiodataId
	user.Biodata.Firstname = req.Firstname
	user.Biodata.Lastname = req.Lastname
	user.Biodata.Email = req.Email
	user.Password = req.Password
	user.RoleId = req.RoleId

	if err = tx.Create(&user).Error; err != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_02",
		}
	}

	otp := models.OtpCode{
		ID:     pkg.GenerateUUID(),
		UserId: req.ID,
		Code:   req.OtpCode,
	}

	if err = tx.Create(&otp).Error; err != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_02",
		}
	}

	return &user, nil
}
