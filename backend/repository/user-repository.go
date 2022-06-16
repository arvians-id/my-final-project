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
	Register(user entity.Users) error
	Login(email string, password string) (entity.Users, error)
	GetUserByID(id int) (entity.Users, error)
	ListUser() ([]entity.Users, error)
	GetLastInsertUser() (entity.Users, error)
	Delete(id int) error
	Update(user entity.Users) error
}

type userRepository struct {
	db *sql.DB
}

// Register is a function to register a new user to the database
func (repository *userRepository) Register(user entity.Users) error {
	var id int

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		return err
	}

	defer database.Close()

	temp, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(temp)

	_, err = database.Exec("INSERT INTO users (name, username, email, password, role, email_verification, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", user.Name, user.Username, user.Email, user.Password, user.Role, user.EmailVerification, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	rows := database.QueryRow("SELECT id FROM users WHERE username = ?", user.Username)

	rows.Scan(&id)

	_, err = database.Exec("INSERT INTO user_details (user_id, phone, gender, type_of_disability, birthdate) VALUES (?, ?, ?, ?, ?)", id, user.Phone, user.Gender, user.DisabilityType, user.Birthdate)

	if err != nil {
		return err
	}

	return nil
}

// Login is a function to login a user by email and password
func (repository *userRepository) Login(email string, password string) (entity.Users, error) {

	var user entity.Users

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		return entity.Users{}, err
	}

	defer database.Close()

	rows := database.QueryRow("SELECT id, name, username, email, password, role FROM users WHERE email = ?", email)

	if err != nil {
		return entity.Users{}, err
	}

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role)

	rows = database.QueryRow("SELECT gender, type_of_disability FROM user_details WHERE user_id = ?", user.Id)

	rows.Scan(&user.Gender, &user.DisabilityType)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return entity.Users{}, err
	}

	return user, err
}

// GetUserByID is a function to get a user by id by database
func (repository *userRepository) GetUserByID(id int) (entity.Users, error) {

	var user entity.Users

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		return entity.Users{}, err
	}

	defer database.Close()

	rows := database.QueryRow("SELECT users.id, users.name, users.username, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.address, user_details.birthdate, user_details.image, user_details.description FROM users INNER JOIN user_details ON user_details.user_id = users.id WHERE users.id = ?", id)

	if err != nil {
		return entity.Users{}, err
	}

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Role, &user.Phone, &user.Gender, &user.DisabilityType, &user.Address, &user.Birthdate, &user.Image, &user.Description)
	return user, nil
}

// GetUser is a function to get all users from the database
func (repository *userRepository) ListUser() ([]entity.Users, error) {
	rows, err := repository.db.Query("SELECT users.id, users.name, users.username, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.address, user_details.birthdate, user_details.image, user_details.description FROM users INNER JOIN user_details ON user_details.user_id = users.id")

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
func (repository *userRepository) GetLastInsertUser() (entity.Users, error) {

	var user entity.Users

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		return entity.Users{}, err
	}

	defer database.Close()

	rows := database.QueryRow("SELECT users.id, users.name, users.username, users.email, users.password, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.birthdate, users.email_verification, users.created_at, users.updated_at FROM users INNER JOIN user_details ON user_details.user_id = users.id WHERE users.id = (SELECT MAX(id) FROM users)")

	if err != nil {
		return entity.Users{}, err
	}

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role, &user.Phone, &user.Gender, &user.DisabilityType, &user.Birthdate, &user.EmailVerification, &user.CreatedAt, &user.UpdatedAt)
	return user, nil
}

// Update is a function to update a user by id to database
func (repository *userRepository) Update(user entity.Users) error {

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	_, err = database.Exec("UPDATE users SET name = ?, username = ?, role = ?, updated_at = ? WHERE id = ?", user.Name, user.Username, user.Role, user.UpdatedAt, user.Id)

	if err != nil {
		panic(err)
	}

	_, err = database.Exec("UPDATE user_details SET phone = ?, gender = ?, type_of_disability = ?, address = ?, birthdate = ?, image = ?, description = ? WHERE user_id = ?", user.Phone, user.Gender, user.DisabilityType, user.Address, user.Birthdate, user.Image, user.Description, user.Id)

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

	rows := database.QueryRow("SELECT name FROM users WHERE id = ?", id)

	rows.Scan(&user.Name)

	if user.Name == "" {
		return fmt.Errorf("User not found")
	}

	_, err = database.Exec("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	_, err = database.Exec("DELETE FROM user_details WHERE user_id = ?", id)

	if err != nil {
		panic(err)
	}

	return nil
}
