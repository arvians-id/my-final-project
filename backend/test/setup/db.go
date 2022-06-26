package setup

import (
	"database/sql"
	"fmt"

	"github.com/rg-km/final-project-engineering-12/backend/config"
)

func SuiteSetup(configuration config.Config) (*sql.DB, error) {
	driver := configuration.Get("DB_CONNECTION")
	databaseName := configuration.Get("DB_DATABASE")
	dsn := fmt.Sprintf("./%v.db", databaseName)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TearDownTest(db *sql.DB) error {
	_, err := db.Exec(`DELETE FROM users;`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM user_details;`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM answers;`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM courses;`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM module_articles;`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM module_submissions;`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM questions;`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM user_course;`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM user_submissions;`)
	if err != nil {
		return err
	}

	return nil
}
