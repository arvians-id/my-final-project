package repository

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"golang.org/x/crypto/bcrypt"
)

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{
		db: database,
	}
}

type UserRepository interface {
	GetUser() []entity.Users
	Insert(user entity.Users)
	Delete(id int) error
	Update(user entity.Users) error
	GetUserByID(id int) entity.Users
	GetLastInsertUser() entity.Users
	Login(email string, password string) entity.Users
}

type userRepository struct {
	db *sql.DB
}

func (repository *userRepository) GetLastInsertUser() entity.Users {

	var user entity.Users

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	rows := database.QueryRow("SELECT id, name, username, email, password, role, email_verification, created_at, updated_at FROM users WHERE id = (SELECT MAX(id) FROM users)")

	if err != nil {
		panic(err)
	}

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role, &user.EmailVerification, &user.CreatedAt, &user.UpdatedAt)
	return user
}

func (repository *userRepository) GetUser() []entity.Users {
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

func (repository *userRepository) Insert(user entity.Users) {

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
}

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

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		panic(err)
	}

	return user
}

func (repository *userRepository) Delete(id int) error {

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	_, err = database.Exec("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	return nil
}

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
