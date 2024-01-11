package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/arioprima/Jobseeker/tree/main/backend/Golang/initializers"
	"github.com/arioprima/Jobseeker/tree/main/backend/Golang/models"
	"github.com/arioprima/Jobseeker/tree/main/backend/Golang/repository"
	"github.com/arioprima/Jobseeker/tree/main/backend/Golang/utils"
	"github.com/go-playground/validator/v10"
	"log"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, request models.LoginInput) (models.LoginResponse, error)
	Register(ctx context.Context, request models.RegisterInput) (string, error)
	VerifyEmail(ctx context.Context, request models.VerifyInput) (models.UserResponse, error)
}

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthServiceImpl(authRepository repository.AuthRepository, db *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{AuthRepository: authRepository, DB: db, Validate: validate}
}

func (auth *AuthServiceImpl) Login(ctx context.Context, request models.LoginInput) (models.LoginResponse, error) {
	//TODO implement me
	tx, err := auth.DB.Begin()
	if err != nil {
		return models.LoginResponse{}, err
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

	user, err := auth.AuthRepository.Login(ctx, tx, request.Email)
	if err != nil || user == nil || !user.IsVerified || !user.IsActive {
		return models.LoginResponse{}, err
	}

	err = utils.VerifyPassword(user.Password, request.Password)

	if err != nil {
		return models.LoginResponse{}, errors.New("invalid password")
	}

	config, _ := initializers.LoadConfig(".")

	tokenPayload := map[string]interface{}{
		"id":         user.UserID,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	}

	token, err := utils.GenerateToken(config.TokenExpiresIn, tokenPayload, config.TokenSecret)
	if err != nil {
		return models.LoginResponse{}, err
	}

	return models.LoginResponse{
		Email:     user.Email,
		FirstUser: user.FirstUser,
		TokenType: "Bearer",
		Token:     token,
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
		return "", fmt.Errorf("kesalahan register user: %v", err)
	}

	utils.SendEmail(&newUser, otp)

	return "check your email", nil
}

func (auth *AuthServiceImpl) VerifyEmail(ctx context.Context, request models.VerifyInput) (models.UserResponse, error) {
	//TODO implement me
	if err := auth.Validate.Struct(request); err != nil {
		return models.UserResponse{}, fmt.Errorf("kesalahan validasi: %v", err)
	}
	tx, err := auth.DB.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return models.UserResponse{}, err
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

	log.Printf("Email: %s, Token: %s", request.Email, request.Token)

	user, err := auth.AuthRepository.VerifyEmail(ctx, tx, request.Token)

	if err != nil {
		return models.UserResponse{}, err
	}

	if user == nil {
		return models.UserResponse{}, errors.New("invalid token")
	}

	if user.Email != request.Email {
		return models.UserResponse{}, errors.New("invalid email")
	}

	if user.VerificationToken != request.Token {
		return models.UserResponse{}, errors.New("invalid token")
	}

	user.IsVerified = true
	user.VerificationToken = ""

	err = auth.AuthRepository.UpdateUserVerificationStatus(ctx, tx, user.Email, user.VerificationToken)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		UserID:    user.UserID,
		Email:     user.Email,
		RoleID:    user.RoleID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}
