package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type EmailVerificationRepository interface {
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.EmailVerification, error)
	FindByEmailAndSignature(ctx context.Context, tx *sql.Tx, verif entity.EmailVerification) (entity.EmailVerification, error)
	Create(ctx context.Context, tx *sql.Tx, verif entity.EmailVerification) (entity.EmailVerification, error)
	Update(ctx context.Context, tx *sql.Tx, verif entity.EmailVerification) (entity.EmailVerification, error)
	Delete(ctx context.Context, tx *sql.Tx, email string) error
}

type emailVerificationRepository struct {
}

func NewEmailVerificationRepository() EmailVerificationRepository {
	return &emailVerificationRepository{}
}

func (repository *emailVerificationRepository) FindByEmailAndSignature(ctx context.Context, tx *sql.Tx, verif entity.EmailVerification) (entity.EmailVerification, error) {
	query := "SELECT * FROM email_verifications WHERE email = ? AND signature = ?"
	rows, err := tx.QueryContext(ctx, query, verif.Email, verif.Signature)
	if err != nil {
		return entity.EmailVerification{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var userRequest entity.EmailVerification
	if rows.Next() {
		err := rows.Scan(&userRequest.Email, &userRequest.Signature, &userRequest.Expired)
		if err != nil {
			return entity.EmailVerification{}, err
		}

		return userRequest, nil
	}

	return userRequest, errors.New("invalid credentials")
}

func (repository *emailVerificationRepository) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.EmailVerification, error) {
	query := "SELECT * FROM email_verifications WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		return entity.EmailVerification{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var userRequest entity.EmailVerification
	if rows.Next() {
		err := rows.Scan(&userRequest.Email, &userRequest.Signature, &userRequest.Expired)
		if err != nil {
			return entity.EmailVerification{}, err
		}

		return userRequest, nil
	}

	return userRequest, nil
}

func (repository *emailVerificationRepository) Create(ctx context.Context, tx *sql.Tx, verif entity.EmailVerification) (entity.EmailVerification, error) {
	query := "INSERT INTO email_verifications (email,signature,expired) VALUES(?,?,?)"
	_, err := tx.ExecContext(ctx, query, verif.Email, verif.Signature, verif.Expired)
	if err != nil {
		return entity.EmailVerification{}, err
	}

	return verif, nil
}

func (repository *emailVerificationRepository) Update(ctx context.Context, tx *sql.Tx, verif entity.EmailVerification) (entity.EmailVerification, error) {
	query := "UPDATE email_verifications SET signature = ?, expired = ? WHERE email = ?"
	_, err := tx.ExecContext(ctx, query, verif.Signature, verif.Expired, verif.Email)
	if err != nil {
		return entity.EmailVerification{}, err
	}

	return verif, nil
}

func (repository *emailVerificationRepository) Delete(ctx context.Context, tx *sql.Tx, email string) error {
	query := "DELETE FROM email_verifications WHERE email = ?"
	_, err := tx.ExecContext(ctx, query, email)
	if err != nil {
		return err
	}

	return nil
}
