package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/arioprima/jobseeker/tree/main/backend/initializers"
	"github.com/arioprima/jobseeker/tree/main/backend/models"
	"github.com/arioprima/jobseeker/tree/main/backend/repository"
	"github.com/arioprima/jobseeker/tree/main/backend/utils"
	"github.com/go-playground/validator/v10"
)

type AuthService interface {
	Login(ctx context.Context, request models.LoginInput) (map[string]interface{}, error)
	Register(ctx context.Context, request models.RegisterInput) (string, error)
	VerifyEmail(ctx context.Context, request models.VerifyInput) (map[string]interface{}, error)
}

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthServiceImpl(authRepository repository.AuthRepository, db *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{AuthRepository: authRepository, DB: db, Validate: validate}
}

func (auth *AuthServiceImpl) Login(ctx context.Context, request models.LoginInput) (map[string]interface{}, error) {
	tx, err := auth.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err := tx.Rollback()
			if err != nil {
				log.Println("Error rolling back transaction:", err)
			}
		} else {
			err := tx.Commit()
			if err != nil {
				log.Println("Error committing transaction:", err)
			}
		}
	}()

	user, err := auth.AuthRepository.Login(ctx, tx, request.Email, request.Password)
	if err != nil {
		return map[string]interface{}{
			"message": err.Error(),
		}, err
	}
	err = utils.VerifyPassword(user.Password, request.Password)

	if err != nil {
		fmt.Println("Error verifying password:", err) // Tambahkan ini
		return map[string]interface{}{
			"message": "error verifikasi password",
		}, err
	}

	config, _ := initializers.LoadConfig(".")

	tokenPayload := map[string]interface{}{
		"id":         user.UserID,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"firs_user":  user.FirstUser,
	}

	token, err := utils.GenerateToken(config.TokenExpiresIn, tokenPayload, config.TokenSecret)
	if err != nil {
		fmt.Println("Error generating token:", err) // Tambahkan ini
		return map[string]interface{}{
			"message": "error generate token",
		}, err
	}

	return map[string]interface{}{
		"token_type": "Bearer",
		"token":      token,
	}, nil
}

func (auth *AuthServiceImpl) Register(ctx context.Context, request models.RegisterInput) (string, error) {
	//TODO implement me
	if err := auth.Validate.Struct(request); err != nil {
		return "", fmt.Errorf("kesalahan validasi: %v", err)
	}

	tx, err := auth.DB.Begin()
	if err != nil {
		return "", fmt.Errorf("kesalahan memulai transaksi: %v", err)
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

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return "", fmt.Errorf("kesalahan hashing password: %v", err)
	}

	now := time.Now()
	newUser := models.User{
		UserID:     utils.GenerateUUID(),
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Email:      request.Email,
		Password:   hashedPassword,
		FirstUser:  true,
		IsActive:   true,
		IsVerified: false,
		RoleID:     request.RoleID,
		CreatedAt:  now,
	}

	otp := utils.GenerateOTP()

	newUser.VerificationToken = otp
	_, err = auth.AuthRepository.Register(ctx, tx, &newUser)
	if err != nil {
		return "email already registered", err
	}

	utils.SendEmail(&newUser, otp)

	return "check your email", nil
}

func (auth *AuthServiceImpl) VerifyEmail(ctx context.Context, request models.VerifyInput) (map[string]interface{}, error) {
	//TODO implement me
	if err := auth.Validate.Struct(request); err != nil {
		log.Printf("Validation error: %v", err)
		return nil, fmt.Errorf("validation error: %v", err)
	}

	tx, err := auth.DB.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("Error rolling back transaction: %v", rollbackErr)
			}
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				log.Printf("Error committing transaction: %v", commitErr)
			}
		}
	}()
	user, err := auth.AuthRepository.VerifyEmail(ctx, tx, request.Token)
	if err != nil {
		return map[string]interface{}{
			"message": fmt.Sprintf("%v", err),
		}, err
	}

	// Perbaiki pemanggilan fungsi UpdateUserVerificationStatus untuk mencocokkan perubahan
	user, err = auth.AuthRepository.UpdateUserVerificationStatus(ctx, tx, request.Email, request.Token)
	if err != nil {
		// log.Printf("Error updating user verification status: %v", err)
		return map[string]interface{}{
			"message": fmt.Sprintf("%v", err),
		}, err
	}

	// log.Println("End VerifyEmail Function")

	return map[string]interface{}{
		"id":         user.UserID,
		"email":      user.Email,
		"role_id":    user.RoleID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"first_user": user.FirstUser,
	}, nil
}
