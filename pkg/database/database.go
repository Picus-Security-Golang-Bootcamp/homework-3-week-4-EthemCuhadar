package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPsqlDB will take database connection information
// and creates a local database server on Postgresql.
// After that, it will connect to database.
func NewPsqlDB() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("BOOKSTORE_DB_HOST"),
		os.Getenv("BOOKSTORE_DB_PORT"),
		os.Getenv("BOOKSTORE_DB_USERNAME"),
		os.Getenv("BOOKSTORE_DB_NAME"),
		os.Getenv("BOOKSTORE_DB_PASSWORD"),
	)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("cannot open database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("connected to local psql server")
	return db, nil
}
