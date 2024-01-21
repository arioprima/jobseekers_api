package service

import (
	"context"
	"database/sql"
	"github.com/arioprima/jobseeker/tree/main/backend/models"
	"github.com/arioprima/jobseeker/tree/main/backend/repository"
	"github.com/arioprima/jobseeker/tree/main/backend/utils"
	"github.com/go-playground/validator/v10"
	"log"
	"time"
)

type AdminService interface {
	Save(ctx context.Context, request models.AdminInput) (map[string]interface{}, error)
	Update(ctx context.Context, request models.AdminInput) (map[string]interface{}, error)
}

type AdminServiceImpl struct {
	AdminRepository repository.AdminRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewAdminServiceImpl(adminRepository repository.AdminRepository, db *sql.DB, validate *validator.Validate) AdminService {
	return &AdminServiceImpl{AdminRepository: adminRepository, DB: db, Validate: validate}
}

func (adminService *AdminServiceImpl) Save(ctx context.Context, request models.AdminInput) (map[string]interface{}, error) {
	//TODO implement me
	tx, err := adminService.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			// Terjadi kesalahan, rollback transaksi
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("Kesalahan rollback transaksi: %v", rollbackErr)
			}
			log.Printf("Panic terjadi: %v", r)
		} else {
			// Tidak ada kesalahan, commit transaksi
			if commitErr := tx.Commit(); commitErr != nil {
				log.Printf("Kesalahan commit transaksi: %v", commitErr)
				// Jika terjadi kesalahan commit, rollback transaksi
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					log.Printf("Kesalahan rollback transaksi setelah kesalahan commit: %v", rollbackErr)
				}
			}
		}
	}()

	now := time.Now()
	formattedNow := now.Format("2006-01-02 15:04:05")

	// Parse waktu yang telah diformat menjadi objek time.Time
	parsedTime, err := time.Parse("2006-01-02 15:04:05", formattedNow)
	if err != nil {
		return nil, err
	}
	newUser := models.AdminUser{
		AdminID:    utils.GenerateUUID(),
		BirthPlace: request.BirthPlace,
		BirthDate:  request.BirthDate,
		Phone:      request.Phone,
		Address:    request.Address,
		UserID:     request.UserID,
		CreatedAt:  parsedTime,
		UpdatedAt:  parsedTime,
	}
	request.FirstUser = false

	user, err := adminService.AdminRepository.Save(ctx, tx, &newUser)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user": user,
	}, nil
}

func (adminService *AdminServiceImpl) Update(ctx context.Context, request models.AdminInput) (map[string]interface{}, error) {
	//TODO implement me
	tx, err := adminService.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			// Terjadi kesalahan, rollback transaksi
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("Kesalahan rollback transaksi: %v", rollbackErr)
			}
			log.Printf("Panic terjadi: %v", r)
		} else {
			// Tidak ada kesalahan, commit transaksi
			if commitErr := tx.Commit(); commitErr != nil {
				log.Printf("Kesalahan commit transaksi: %v", commitErr)
				// Jika terjadi kesalahan commit, rollback transaksi
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					log.Printf("Kesalahan rollback transaksi setelah kesalahan commit: %v", rollbackErr)
				}
			}
		}
	}()
	//SQL untuk mendapatkan id admin
	SQL := "select id from admin where user_id = ?"
	row := tx.QueryRowContext(ctx, SQL, request.UserID)

	var adminID string
	err = row.Scan(&adminID)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	formattedNow := now.Format("2006-01-02 15:04:05")

	// Parse waktu yang telah diformat menjadi objek time.Time
	parsedTime, err := time.Parse("2006-01-02 15:04:05", formattedNow)
	if err != nil {
		return nil, err
	}
	newUser := models.AdminUser{
		AdminID:    adminID,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		BirthPlace: request.BirthPlace,
		BirthDate:  request.BirthDate,
		Phone:      request.Phone,
		Address:    request.Address,
		UserID:     request.UserID,
		UpdatedAt:  parsedTime,
	}

	user, err := adminService.AdminRepository.Update(ctx, tx, &newUser)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user": user,
	}, nil
}
