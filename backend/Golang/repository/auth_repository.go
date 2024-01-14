package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arioprima/Jobseeker/tree/main/backend/Golang/models"
	"log"
)

type AuthRepository interface {
	Login(ctx context.Context, tx *sql.Tx, email string) (*models.User, error)
	Register(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error)
	VerifyEmail(ctx context.Context, tx *sql.Tx, otpCode string) (*models.User, error)
	UpdateUserVerificationStatus(ctx context.Context, tx *sql.Tx, email, token string) error
}

type authRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepositoryImpl(db *sql.DB) AuthRepository {
	return &authRepositoryImpl{DB: db}
}

func (auth *authRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, email string) (*models.User, error) {
	//TODO implement me
	SQL := `SELECT * FROM users WHERE email = ? and is_verified = true and is_active = true`
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
			return nil, nil
		}
		return nil, err
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
			log.Printf("Kesalahan memeriksa email yang sudah ada: %v", err)
			return nil, err
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

func (auth *authRepositoryImpl) UpdateUserVerificationStatus(ctx context.Context, tx *sql.Tx, email, token string) error {
	//TODO implement me
	SQL := `UPDATE users SET is_verified = true WHERE email = ? and verification_token = ?`
	result, err := tx.ExecContext(ctx, SQL, email, token)
	if err != nil {
		log.Printf("Error updating user verification status: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}

	log.Printf("Rows affected: %d", rowsAffected)

	if rowsAffected == 0 {
		log.Println("No rows updated")
		return errors.New("no rows updated")
	}

	return nil

}
