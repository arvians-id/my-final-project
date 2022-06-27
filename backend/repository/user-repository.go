package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user entity.Users) error
	Login(ctx context.Context, tx *sql.Tx, data model.GetUserLogin) (entity.Users, error)
	UpdateRole(ctx context.Context, tx *sql.Tx, id int) (entity.Users, error)
	GetUserByID(ctx context.Context, tx *sql.Tx, id int) (entity.Users, error)
	ListUser(ctx context.Context, tx *sql.Tx) ([]entity.Users, error)
	GetLastInsertUser(ctx context.Context, tx *sql.Tx) (entity.Users, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	Update(ctx context.Context, tx *sql.Tx, user entity.Users) error
	CheckUserByEmail(ctx context.Context, tx *sql.Tx, email string) error
	UpdateVerifiedAt(ctx context.Context, tx *sql.Tx, timeVerifiedAt time.Time, email string) error
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// Register is a function to register a new user to the database
func (repository *userRepository) Register(ctx context.Context, tx *sql.Tx, user entity.Users) error {
	var id int
	var email, username string
	var emailArr, usernameArr []string

	rowsCheck, err := tx.QueryContext(ctx, "SELECT email, username FROM users")

	if err != nil {
		return err
	}

	for rowsCheck.Next() {
		rowsCheck.Scan(&email, &username)
		emailArr = append(emailArr, email)
		usernameArr = append(usernameArr, username)
	}

	for _, value := range usernameArr {
		if value == user.Username {
			return fmt.Errorf("username has been registered")
		}
	}

	for _, value := range emailArr {
		if value == user.Email {
			return fmt.Errorf("email has been registered")
		}
	}

	temp, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(temp)

	_, err = tx.ExecContext(ctx, "INSERT INTO users (name, username, email, password, role, email_verification, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", user.Name, user.Username, user.Email, user.Password, user.Role, user.EmailVerification, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	rows := tx.QueryRowContext(ctx, "SELECT id FROM users WHERE username = ?", user.Username)

	rows.Scan(&id)

	_, err = tx.ExecContext(ctx, "INSERT INTO user_details (user_id, phone, gender, type_of_disability, birthdate) VALUES (?, ?, ?, ?, ?)", id, user.Phone, user.Gender, user.DisabilityType, user.Birthdate)

	if err != nil {
		return err
	}

	return nil
}

// Login is a function to login a user by email and password
func (repository *userRepository) Login(ctx context.Context, tx *sql.Tx, data model.GetUserLogin) (entity.Users, error) {

	var user entity.Users

	rows := tx.QueryRowContext(ctx, "SELECT id, name, username, email, password, role FROM users WHERE email = ?", data.Email)

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role)

	rows = tx.QueryRowContext(ctx, "SELECT gender, type_of_disability FROM user_details WHERE user_id = ?", user.Id)

	rows.Scan(&user.Gender, &user.DisabilityType)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return entity.Users{}, err
	}

	return user, err
}

// UpdateRole is a function to update a user's role by database
func (repository *userRepository) UpdateRole(ctx context.Context, tx *sql.Tx, id int) (entity.Users, error) {
	var user entity.Users

	_, err := tx.ExecContext(ctx, "UPDATE users SET role = ? WHERE id = ?", 1, id)

	if err != nil {
		return entity.Users{}, err
	}

	rows := tx.QueryRowContext(ctx, "SELECT users.id, users.name, users.username, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.address, user_details.birthdate, user_details.image, user_details.description FROM users INNER JOIN user_details ON user_details.user_id = users.id WHERE users.id = ?", id)

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Role, &user.Phone, &user.Gender, &user.DisabilityType, &user.Address, &user.Birthdate, &user.Image, &user.Description)

	return user, nil
}

// GetUserByID is a function to get a user by id by database
func (repository *userRepository) GetUserByID(ctx context.Context, tx *sql.Tx, id int) (entity.Users, error) {

	var user entity.Users

	rows := tx.QueryRowContext(ctx, "SELECT users.id, users.name, users.username, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.address, user_details.birthdate, user_details.image, user_details.description FROM users INNER JOIN user_details ON user_details.user_id = users.id WHERE users.id = ?", id)

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Role, &user.Phone, &user.Gender, &user.DisabilityType, &user.Address, &user.Birthdate, &user.Image, &user.Description)
	return user, nil
}

// GetUser is a function to get all users from the database
func (repository *userRepository) ListUser(ctx context.Context, tx *sql.Tx) ([]entity.Users, error) {
	rows, err := tx.QueryContext(ctx, "SELECT users.id, users.name, users.username, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.address, user_details.birthdate, user_details.image, user_details.description FROM users INNER JOIN user_details ON user_details.user_id = users.id")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []entity.Users
	for rows.Next() {
		var user entity.Users

		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Role, &user.Phone, &user.Gender, &user.DisabilityType, &user.Address, &user.Birthdate, &user.Image, &user.Description)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetLastInsertUser is a function to get the last inserted user from the database
func (repository *userRepository) GetLastInsertUser(ctx context.Context, tx *sql.Tx) (entity.Users, error) {

	var user entity.Users

	rows := tx.QueryRowContext(ctx, "SELECT users.id, users.name, users.username, users.email, users.password, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.birthdate, users.email_verification, users.created_at, users.updated_at FROM users INNER JOIN user_details ON user_details.user_id = users.id WHERE users.id = (SELECT MAX(id) FROM users)")

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role, &user.Phone, &user.Gender, &user.DisabilityType, &user.Birthdate, &user.EmailVerification, &user.CreatedAt, &user.UpdatedAt)
	return user, nil
}

// Update is a function to update a user by id to database
func (repository *userRepository) Update(ctx context.Context, tx *sql.Tx, user entity.Users) error {
	var email, username string
	var emailArr, usernameArr []string

	rowsCheck, err := tx.QueryContext(ctx, "SELECT email, username FROM users")

	if err != nil {
		return err
	}

	for rowsCheck.Next() {
		rowsCheck.Scan(&email, &username)
		emailArr = append(emailArr, email)
		usernameArr = append(usernameArr, username)
	}

	for _, value := range usernameArr {
		if value == user.Username {
			return fmt.Errorf("username has been registered")
		}
	}

	for _, value := range emailArr {
		if value == user.Email {
			return fmt.Errorf("email has been registered")
		}
	}

	_, err = tx.ExecContext(ctx, "UPDATE users SET name = ?, username = ?, role = ?, updated_at = ? WHERE id = ?", user.Name, user.Username, user.Role, user.UpdatedAt, user.Id)

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE user_details SET phone = ?, gender = ?, type_of_disability = ?, address = ?, birthdate = ?, image = ?, description = ? WHERE user_id = ?", user.Phone, user.Gender, user.DisabilityType, user.Address, user.Birthdate, user.Image, user.Description, user.Id)

	if err != nil {
		return err
	}

	return nil
}

// Delete is a function to delete a user by id to database
func (repository *userRepository) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	var user entity.Users

	rows := tx.QueryRowContext(ctx, "SELECT name FROM users WHERE id = ?", id)

	rows.Scan(&user.Name)

	if user.Name == "" {
		return fmt.Errorf("user not found")
	}

	_, err := tx.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM user_details WHERE user_id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func (repository *userRepository) CheckUserByEmail(ctx context.Context, tx *sql.Tx, email string) error {
	query := "SELECT * FROM users WHERE email = ?"
	queryContext, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		return err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	if queryContext.Next() {
		return nil
	}

	return errors.New("the user with the email was not found")
}

func (repository *userRepository) UpdateVerifiedAt(ctx context.Context, tx *sql.Tx, timeVerifiedAt time.Time, email string) error {
	query := "UPDATE users SET email_verification = ? WHERE email = ?"
	_, err := tx.ExecContext(ctx, query, timeVerifiedAt, email)
	if err != nil {
		return err
	}

	return nil
}
