package config

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func NewSQLite(configuration Config) *sql.DB {
	// Setup DB
	driver := configuration.Get("DB_CONNECTION")
	databaseName := configuration.Get("DB_DATABASE")

	dsn := fmt.Sprintf("./%v.db", databaseName)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Limit connection with db pooling
	setMaxIdleConns, err := strconv.Atoi(configuration.Get("SQLITE_POOL_MIN"))
	if err != nil {
		panic(err)
	}
	setMaxOpenConns, err := strconv.Atoi(configuration.Get("SQLITE_POOL_MAX"))
	if err != nil {
		panic(err)
	}
	setConnMaxIdleTime, err := strconv.Atoi(configuration.Get("SQLITE_MAX_IDLE_TIME_SECOND"))
	if err != nil {
		panic(err)
	}
	setConnMaxLifetime, err := strconv.Atoi(configuration.Get("SQLITE_MAX_LIFE_TIME_SECOND"))
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(setMaxIdleConns)                                    // minimal connection
	db.SetMaxOpenConns(setMaxOpenConns)                                    // maximal connection
	db.SetConnMaxLifetime(time.Duration(setConnMaxIdleTime) * time.Minute) // unused connections will be deleted
	db.SetConnMaxIdleTime(time.Duration(setConnMaxLifetime) * time.Minute) // connection that can be used

	return db
}
