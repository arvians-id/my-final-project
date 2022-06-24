package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
	"time"
)

type EmailService interface {
	SendEmailWithText(toEmail string, message string) error
	VerifyEmail(ctx context.Context, request model.GetEmailVerificationRequest) error
}

type emailService struct {
	EmailVerificationRepository repository.EmailVerificationRepository
	UserRepository              repository.UserRepository
	DB                          *sql.DB
}

func NewEmailService(verificationRepository *repository.EmailVerificationRepository, userRepository *repository.UserRepository, db *sql.DB) EmailService {
	return &emailService{
		EmailVerificationRepository: *verificationRepository,
		UserRepository:              *userRepository,
		DB:                          db,
	}
}

func (service *emailService) SendEmailWithText(toEmail string, message string) error {
	go func() {
		mailer := gomail.NewMessage()
		mailer.SetHeader("From", os.Getenv("MAIL_FROM_ADDRESS"))
		mailer.SetHeader("To", toEmail)
		mailer.SetHeader("Subject", "Test Bang")
		mailer.SetBody("text/html", message)

		port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
		if err != nil {
			panic(err)
		}

		dialer := gomail.NewDialer(
			os.Getenv("MAIL_HOST"),
			port,
			os.Getenv("MAIL_USERNAME"),
			os.Getenv("MAIL_PASSWORD"),
		)

		err = dialer.DialAndSend(mailer)
		if err != nil {
			panic(err)
		}
	}()

	return nil
}

func (service *emailService) VerifyEmail(ctx context.Context, request model.GetEmailVerificationRequest) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	// Check if email and token is exists in table email verifications
	emailVerification := entity.EmailVerification{
		Email:     request.Email,
		Signature: request.Signature,
	}

	dataEmailVerification, err := service.EmailVerificationRepository.FindByEmailAndSignature(ctx, tx, emailVerification)
	if err != nil {
		return err
	}

	// Check token expired
	now := time.Now()

	// if expired
	if now.Unix() > int64(dataEmailVerification.Expired) {
		err := service.EmailVerificationRepository.Delete(ctx, tx, dataEmailVerification.Email)
		if err != nil {
			return err
		}
		return errors.New("email verification is expired")
	}

	// if not
	// Be sure the email is exists in users table
	err = service.UserRepository.CheckUserByEmail(ctx, tx, dataEmailVerification.Email)
	if err != nil {
		return err
	}

	// Update the email verified at
	err = service.UserRepository.UpdateVerifiedAt(ctx, tx, utils.TimeNow(), dataEmailVerification.Email)
	if err != nil {
		return err
	}

	// Delete data from table password_reset
	err = service.EmailVerificationRepository.Delete(ctx, tx, dataEmailVerification.Email)
	if err != nil {
		return err
	}

	return nil
}
