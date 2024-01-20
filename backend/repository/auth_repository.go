package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/arioprima/jobseeker/tree/main/backend/models"
	"github.com/arioprima/jobseeker/tree/main/backend/utils"
)

type AuthRepository interface {
	Login(ctx context.Context, tx *sql.Tx, email, password string) (*models.User, error)
	Register(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error)
	VerifyEmail(ctx context.Context, tx *sql.Tx, otpCode string) (*models.User, error)
	UpdateUserVerificationStatus(ctx context.Context, tx *sql.Tx, email, token string) (*models.User, error)
}

type authRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepositoryImpl(db *sql.DB) AuthRepository {
	return &authRepositoryImpl{DB: db}
}

func (auth *authRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, email, password string) (*models.User, error) {
	//TODO implement me
	SQL := `SELECT * FROM users WHERE email = ?`
	row := tx.QueryRowContext(ctx, SQL, email)

	var user models.User
	err := row.Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.FirstUser,
		&user.IsActive,
		&user.IsVerified,
		&user.VerificationToken,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("email is not registered")
		}
		return nil, err
	}

	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		// log.Printf("Error verifying password: %v", err)
		return nil, errors.New("password is wrong")
	}

	if !user.IsActive {
		return nil, errors.New("your account is not active")
	}

	if !user.IsVerified {
		return nil, errors.New("your account is not verified")
	}

	if !user.FirstUser {
		return nil, errors.New("your account is not first user")
	}

	return &user, nil
}

func (auth *authRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error) {
	//TODO implement me
	SQL := `SELECT * FROM users WHERE email = ?`
	row := tx.QueryRowContext(ctx, SQL, user.Email)

	var existingUser models.User
	err := row.Scan(
		&existingUser.UserID,
		&existingUser.FirstName,
		&existingUser.LastName,
		&existingUser.Email,
		&existingUser.Password,
		&existingUser.FirstUser,
		&existingUser.IsActive,
		&existingUser.IsVerified,
		&existingUser.VerificationToken,
		&existingUser.RoleID,
		&existingUser.CreatedAt,
	)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			// log.Printf("Kesalahan memeriksa email yang sudah ada: %v", err)
			return nil, errors.New("email already registered")
		}
	}

	if existingUser.Email != "" {
		return nil, errors.New("email sudah terdaftar")
	}

	SQL = `INSERT INTO users (id, first_name, last_name, email, password, first_user, is_active, is_verified, verification_token, role_id, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, SQL, user.UserID, user.FirstName, user.LastName, user.Email, user.Password, user.FirstUser, user.IsActive, user.IsVerified, user.VerificationToken, user.RoleID, user.CreatedAt)
	if err != nil {
		log.Printf("Kesalahan saat insert data user: %v", err)
		return nil, err
	}

	return user, nil
}

func (auth *authRepositoryImpl) VerifyEmail(ctx context.Context, tx *sql.Tx, otpCode string) (*models.User, error) {
	//TODO implement me
	SQL := `SELECT * FROM users WHERE verification_token = ?`
	row := tx.QueryRowContext(ctx, SQL, otpCode)

	var user models.User
	err := row.Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.FirstUser,
		&user.IsActive,
		&user.IsVerified,
		&user.VerificationToken,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
func (auth *authRepositoryImpl) UpdateUserVerificationStatus(ctx context.Context, tx *sql.Tx, email, token string) (*models.User, error) {
	// TODO implement me
	SQL := `UPDATE users SET is_verified = true WHERE email = ? and verification_token = ?`
	result, err := tx.ExecContext(ctx, SQL, email, token)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		var user models.User

		// Pindahkan pemindaian baris ke sini setelah memastikan bahwa token valid
		row := tx.QueryRowContext(ctx, "SELECT * FROM users WHERE email = ? FOR UPDATE", email)
		err = row.Scan(
			&user.UserID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.FirstUser,
			&user.IsActive,
			&user.IsVerified,
			&user.VerificationToken,
			&user.RoleID,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("email is not registered: %v", err)
				return nil, errors.New("email is not registered")
			}
			return nil, err
		} else if user.IsVerified {
			log.Printf("email is already verified: %v", err)
			return nil, errors.New("email is already verified")
		}

		log.Printf("token is not valid: %v", err)
		return nil, errors.New("token is not valid")
	}

	var user models.User
	row := tx.QueryRowContext(ctx, "SELECT * FROM users WHERE email = ? FOR UPDATE", email)
	err = row.Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.FirstUser,
		&user.IsActive,
		&user.IsVerified,
		&user.VerificationToken,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return &user, nil
}
