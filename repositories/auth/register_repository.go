package repositories

import (
	"context"
	"github.com/arioprima/jobseekers_api/models"
	"github.com/arioprima/jobseekers_api/schemas"
	"gorm.io/gorm"
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

	checkUserAccount := tx.Where("email = ?", req.Email).First(&user)
	if checkUserAccount.RowsAffected > 0 {
		return nil, &schemas.SchemaDatabaseError{
			Code: 400,
			Type: "error_01",
		}
	}

	user.Biodata.Firstname = req.Firstname
	user.Biodata.Lastname = req.Lastname
	user.Biodata.Email = req.Email
	user.Password = req.Password

	if err = tx.Create(&user).Error; err != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: 500,
			Type: "error_02",
		}
	}
	return &user, nil
}
