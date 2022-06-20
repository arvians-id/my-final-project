package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
<<<<<<< Updated upstream
	Register(user entity.Users) error
	Login(email string, password string) (entity.Users, error)
	GetUserByID(id int) (entity.Users, error)
	ListUser() ([]entity.Users, error)
	GetLastInsertUser() (entity.Users, error)
=======
<<<<<<< Updated upstream
	Register(user entity.Users)
	Login(email string, password string) entity.Users
	GetUserByID(id int) entity.Users
	ListUser() []entity.Users
	GetLastInsertUser() entity.Users
>>>>>>> Stashed changes
	Delete(id int) error
	Update(user entity.Users) error
=======
	Register(ctx context.Context, tx *sql.Tx, user entity.Users) error
	Login(ctx context.Context, tx *sql.Tx, data model.GetUserLogin) (entity.Users, error)
	UpdateRole(ctx context.Context, tx *sql.Tx, id int) (entity.Users, error)
	GetUserByID(ctx context.Context, tx *sql.Tx, id int) (entity.Users, error)
	ListUser(ctx context.Context, tx *sql.Tx) ([]entity.Users, error)
	GetLastInsertUser(ctx context.Context, tx *sql.Tx) (entity.Users, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	Update(ctx context.Context, tx *sql.Tx, user entity.Users) error
>>>>>>> Stashed changes
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// Register is a function to register a new user to the database
<<<<<<< Updated upstream
func (repository *userRepository) Register(user entity.Users) error {
=======
<<<<<<< Updated upstream
func (repository *userRepository) Register(user entity.Users) {
>>>>>>> Stashed changes
	var id int

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		return err
	}

	defer database.Close()

=======
func (repository *userRepository) Register(ctx context.Context, tx *sql.Tx, user entity.Users) error {
	var id int

>>>>>>> Stashed changes
	temp, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(temp)

	_, err := tx.ExecContext(ctx, "INSERT INTO users (name, username, email, password, role, email_verification, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", user.Name, user.Username, user.Email, user.Password, user.Role, user.EmailVerification, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	rows := tx.QueryRowContext(ctx, "SELECT id FROM users WHERE username = ?", user.Username)

	rows.Scan(&id)

<<<<<<< Updated upstream
	_, err = database.Exec("INSERT INTO user_details (user_id, phone, gender, type_of_disability, birthdate) VALUES (?, ?, ?, ?, ?)", id, user.Phone, user.Gender, user.DisabilityType, user.Birthdate)
=======
<<<<<<< Updated upstream
	_, err = database.Exec("INSERT INTO user_details (user_id, gender, type_of_disability, birthdate) VALUES (?, ?, ?, ?)", id, user.Gender, user.DisabilityType, user.Birthdate)
=======
	_, err = tx.ExecContext(ctx, "INSERT INTO user_details (user_id, phone, gender, type_of_disability, birthdate) VALUES (?, ?, ?, ?, ?)", id, user.Phone, user.Gender, user.DisabilityType, user.Birthdate)
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	if err != nil {
		return err
	}

	return nil
}

// Login is a function to login a user by email and password
<<<<<<< Updated upstream
func (repository *userRepository) Login(email string, password string) (entity.Users, error) {
=======
<<<<<<< Updated upstream
func (repository *userRepository) Login(email string, password string) entity.Users {
>>>>>>> Stashed changes

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
=======
func (repository *userRepository) Login(ctx context.Context, tx *sql.Tx, data model.GetUserLogin) (entity.Users, error) {

	var user entity.Users

	rows := tx.QueryRowContext(ctx, "SELECT id, name, username, email, password, role FROM users WHERE email = ?", data.Email)
>>>>>>> Stashed changes

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role)

<<<<<<< Updated upstream
	rows = database.QueryRow("SELECT gender, type_of_disability FROM user_details WHERE user_id = ?", user.Id)
=======
<<<<<<< Updated upstream
	rows = database.QueryRow("SELECT gender, type_of_disability, birthdate FROM user_details WHERE user_id = ?", user.Id)
=======
	rows = tx.QueryRowContext(ctx, "SELECT gender, type_of_disability FROM user_details WHERE user_id = ?", user.Id)
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	rows.Scan(&user.Gender, &user.DisabilityType)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return entity.Users{}, err
	}

	return user, err
}

<<<<<<< Updated upstream
// GetUserByID is a function to get a user by id by database
func (repository *userRepository) GetUserByID(id int) (entity.Users, error) {

=======
// UpdateRole is a function to update a user's role by database
func (repository *userRepository) UpdateRole(ctx context.Context, tx *sql.Tx, id int) (entity.Users, error) {
>>>>>>> Stashed changes
	var user entity.Users

	_, err := tx.ExecContext(ctx, "UPDATE users SET role = ? WHERE id = ?", 1, id)

	if err != nil {
		return entity.Users{}, err
	}

	rows := tx.QueryRowContext(ctx, "SELECT users.id, users.name, users.username, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.address, user_details.birthdate, user_details.image, user_details.description FROM users INNER JOIN user_details ON user_details.user_id = users.id WHERE users.id = ?", id)

<<<<<<< Updated upstream
	rows := database.QueryRow("SELECT users.id, users.name, users.username, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.address, user_details.birthdate, user_details.image, user_details.description FROM users INNER JOIN user_details ON user_details.user_id = users.id WHERE users.id = ?", id)
=======
<<<<<<< Updated upstream
	rows := database.QueryRow("SELECT id, name, username, email, password, role, email_verification, created_at, updated_at FROM users WHERE id = ?", id)
>>>>>>> Stashed changes

	if err != nil {
		return entity.Users{}, err
	}
=======
	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Role, &user.Phone, &user.Gender, &user.DisabilityType, &user.Address, &user.Birthdate, &user.Image, &user.Description)

	return user, nil
}

// GetUserByID is a function to get a user by id by database
func (repository *userRepository) GetUserByID(ctx context.Context, tx *sql.Tx, id int) (entity.Users, error) {

	var user entity.Users

	rows := tx.QueryRowContext(ctx, "SELECT users.id, users.name, users.username, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.address, user_details.birthdate, user_details.image, user_details.description FROM users INNER JOIN user_details ON user_details.user_id = users.id WHERE users.id = ?", id)
>>>>>>> Stashed changes

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Role, &user.Phone, &user.Gender, &user.DisabilityType, &user.Address, &user.Birthdate, &user.Image, &user.Description)
	return user, nil
}

// GetUser is a function to get all users from the database
<<<<<<< Updated upstream
func (repository *userRepository) ListUser() ([]entity.Users, error) {
	rows, err := repository.db.Query("SELECT users.id, users.name, users.username, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.address, user_details.birthdate, user_details.image, user_details.description FROM users INNER JOIN user_details ON user_details.user_id = users.id")
=======
<<<<<<< Updated upstream
func (repository *userRepository) ListUser() []entity.Users {
	rows, err := repository.db.Query("SELECT id, name, username, email, password, role, email_verification FROM users")
=======
func (repository *userRepository) ListUser(ctx context.Context, tx *sql.Tx) ([]entity.Users, error) {
	rows, err := tx.QueryContext(ctx, "SELECT users.id, users.name, users.username, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.address, user_details.birthdate, user_details.image, user_details.description FROM users INNER JOIN user_details ON user_details.user_id = users.id")
>>>>>>> Stashed changes
>>>>>>> Stashed changes

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
<<<<<<< Updated upstream
func (repository *userRepository) GetLastInsertUser() (entity.Users, error) {
=======
<<<<<<< Updated upstream
func (repository *userRepository) GetLastInsertUser() entity.Users {
>>>>>>> Stashed changes

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
=======
func (repository *userRepository) GetLastInsertUser(ctx context.Context, tx *sql.Tx) (entity.Users, error) {

	var user entity.Users

	rows := tx.QueryRowContext(ctx, "SELECT users.id, users.name, users.username, users.email, users.password, users.role, user_details.phone, user_details.gender, user_details.type_of_disability, user_details.birthdate, users.email_verification, users.created_at, users.updated_at FROM users INNER JOIN user_details ON user_details.user_id = users.id WHERE users.id = (SELECT MAX(id) FROM users)")
>>>>>>> Stashed changes

	rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role, &user.Phone, &user.Gender, &user.DisabilityType, &user.Birthdate, &user.EmailVerification, &user.CreatedAt, &user.UpdatedAt)
	return user, nil
}

// Update is a function to update a user by id to database
func (repository *userRepository) Update(ctx context.Context, tx *sql.Tx, user entity.Users) error {

	_, err := tx.ExecContext(ctx, "UPDATE users SET name = ?, username = ?, role = ?, updated_at = ? WHERE id = ?", user.Name, user.Username, user.Role, user.UpdatedAt, user.Id)

	if err != nil {
<<<<<<< Updated upstream
		panic(err)
	}

	defer database.Close()

<<<<<<< Updated upstream
	_, err = database.Exec("UPDATE users SET name = ?, username = ?, role = ?, updated_at = ? WHERE id = ?", user.Name, user.Username, user.Role, user.UpdatedAt, user.Id)

	if err != nil {
		panic(err)
	}

	_, err = database.Exec("UPDATE user_details SET phone = ?, gender = ?, type_of_disability = ?, address = ?, birthdate = ?, image = ?, description = ? WHERE user_id = ?", user.Phone, user.Gender, user.DisabilityType, user.Address, user.Birthdate, user.Image, user.Description, user.Id)
=======
	_, err = database.Exec("UPDATE users SET name = ?, username = ?, email = ?, password = ?, role = ?, email_verification = ?, created_at = ?, updated_at = ? WHERE id = ?", user.Name, user.Username, user.Email, user.Password, user.Role, user.EmailVerification, user.CreatedAt, user.UpdatedAt, user.Id)
=======
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE user_details SET phone = ?, gender = ?, type_of_disability = ?, address = ?, birthdate = ?, image = ?, description = ? WHERE user_id = ?", user.Phone, user.Gender, user.DisabilityType, user.Address, user.Birthdate, user.Image, user.Description, user.Id)
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	if err != nil {
		return err
	}

	return nil
}

// Delete is a function to delete a user by id to database
func (repository *userRepository) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	var user entity.Users

<<<<<<< Updated upstream
	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

<<<<<<< Updated upstream
	rows := database.QueryRow("SELECT name FROM users WHERE id = ?", id)
=======
	rows := database.QueryRow("SELECT id, name, username, email, password, role, email_verification, created_at, updated_at FROM users WHERE id = ?", id)
=======
	rows := tx.QueryRowContext(ctx, "SELECT name FROM users WHERE id = ?", id)
>>>>>>> Stashed changes
>>>>>>> Stashed changes

	rows.Scan(&user.Name)

	if user.Name == "" {
		return fmt.Errorf("user not found")
	}

	_, err := tx.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)

	if err != nil {
		return err
	}

<<<<<<< Updated upstream
=======
	_, err = tx.ExecContext(ctx, "DELETE FROM user_details WHERE user_id = ?", id)

	if err != nil {
		return err
	}

<<<<<<< Updated upstream
	_, err = database.Exec("DELETE FROM user_details WHERE user_id = ?", id)

	if err != nil {
		panic(err)
	}

=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	return nil
}
