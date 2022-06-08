package repository

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{
		db: database,
	}
}

type UserRepository interface {
	FindUser() []entity.User
	Insert(user entity.User)
}

type userRepository struct {
	db *sql.DB
}

func (repository *userRepository) FindUser() []entity.User {
	rows, err := repository.db.Query("SELECT id, name, unique_number, phone, email, password, role, image FROM users")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User

		err := rows.Scan(&user.ID, &user.Name, &user.Unique_number, &user.Phone, &user.Email, &user.Password, &user.Role, &user.Image)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users
}

func (repository *userRepository) Insert(user entity.User) {

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	_, err = database.Exec("INSERT INTO users (id, name, unique_number, phone, email, password, role, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", user.ID, user.Name, user.Unique_number, user.Phone, user.Email, user.Password, user.Role, user.Image, user.Created_at, user.Updated_at)

	if err != nil {
		panic(err)
	}
}
