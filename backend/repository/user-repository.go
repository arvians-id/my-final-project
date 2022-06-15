package repository

import (
	"database/sql"
	"fmt"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"golang.org/x/crypto/bcrypt"
)

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{
		db: database,
	}
}

type UserRepository interface {
	Register(user entity.Users)
	Login(email string, password string) entity.Users
	GetUserByID(id int) entity.Users
	ListUser() []entity.Users
	GetLastInsertUser() entity.Users
	Delete(id int) error
	Update(user entity.Users) error
}

type userRepository struct {
	db *sql.DB
}

// Register is a function to register a new user to the database
func (repository *userRepository) Register(user entity.Users) {
	var id int

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	temp, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(temp)

	_, err = database.Exec("INSERT INTO users (name, username, email, password, role, email_verification, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", user.Name, user.Username, user.Email, user.Password, user.Role, user.EmailVerification, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		panic(err)
	}

	rows := database.QueryRow("SELECT id FROM users WHERE username = ?", user.Username)

	rows.Scan(&id)

	_, err = database.Exec("INSERT INTO user_details (user_id, gender, type_of_disability, birthdate) VALUES (?, ?, ?, ?)", id, user.Gender, user.DisabilityType, user.Birthdate)

	if err != nil {
		panic(err)
	}
}

// Login is a function to login a user by email and password
func (repository *userRepository) Login(email string, password string) entity.Users {

	var user entity.Users

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	rows := database.QueryRow("SELECT id, name, username, email, password, role, email_verification, created_at, updated_at FROM users WHERE email = ?", email)

	if err != nil {
		panic(err)
	}

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role, &user.EmailVerification, &user.CreatedAt, &user.UpdatedAt)

	rows = database.QueryRow("SELECT gender, type_of_disability, birthdate FROM user_details WHERE user_id = ?", user.Id)

	rows.Scan(&user.Gender, &user.DisabilityType, &user.Birthdate)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		panic(err)
	}

	return user
}

// GetUserByID is a function to get a user by id by database
func (repository *userRepository) GetUserByID(id int) entity.Users {

	var user entity.Users

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	rows := database.QueryRow("SELECT id, name, username, email, password, role, email_verification, created_at, updated_at FROM users WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role, &user.EmailVerification, &user.CreatedAt, &user.UpdatedAt)
	return user
}

// GetUser is a function to get all users from the database
func (repository *userRepository) ListUser() []entity.Users {
	rows, err := repository.db.Query("SELECT id, name, username, email, password, role, email_verification FROM users")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var users []entity.Users
	for rows.Next() {
		var user entity.Users

		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role, &user.EmailVerification)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users
}

// GetLastInsertUser is a function to get the last inserted user from the database
func (repository *userRepository) GetLastInsertUser() entity.Users {

	var user entity.Users

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	rows := database.QueryRow("SELECT users.id, users.name, users.username, users.email, users.password, users.role, user_details.gender, user_details.type_of_disability, user_details.birthdate, users.email_verification, users.created_at, users.updated_at FROM users INNER JOIN user_details ON user_details.user_id = users.id WHERE users.id = (SELECT MAX(id) FROM users)")

	if err != nil {
		panic(err)
	}

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role, &user.Gender, &user.DisabilityType, &user.Birthdate, &user.EmailVerification, &user.CreatedAt, &user.UpdatedAt)
	return user
}

// Update is a function to update a user by id to database
func (repository *userRepository) Update(user entity.Users) error {

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	_, err = database.Exec("UPDATE users SET name = ?, username = ?, email = ?, password = ?, role = ?, email_verification = ?, created_at = ?, updated_at = ? WHERE id = ?", user.Name, user.Username, user.Email, user.Password, user.Role, user.EmailVerification, user.CreatedAt, user.UpdatedAt, user.Id)

	if err != nil {
		panic(err)
	}

	return nil
}

// Delete is a function to delete a user by id to database
func (repository *userRepository) Delete(id int) error {
	var user entity.Users

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	rows := database.QueryRow("SELECT id, name, username, email, password, role, email_verification, created_at, updated_at FROM users WHERE id = ?", id)

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role, &user.EmailVerification, &user.CreatedAt, &user.UpdatedAt)

	if user.Name == "" {
		return fmt.Errorf("User not found")
	}

	_, err = database.Exec("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	return nil
}
