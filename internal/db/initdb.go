package db

import (
	"fmt"
	"os"

	"github.com/glebpepega/new1/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func New() (*DB, error) {
	db := &DB{}

	dbInstance, err := gorm.Open(postgres.Open(os.Getenv("DB_CONN_INFO")), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db openning: %w", err)
	}

	db.DB = dbInstance

	if err = db.constructTables(); err != nil {
		return nil, fmt.Errorf("construct tables: %w", err)
	}

	return db, nil
}

func (db *DB) constructTables() error {
	if err := db.DB.AutoMigrate(&models.News{}); err != nil {
		return fmt.Errorf("automigration db news model: %w", err)
	}

	if err := db.DB.AutoMigrate(&models.NewsCategories{}); err != nil {
		return fmt.Errorf("automigration db news categories: %w", err)
	}

	return nil
}
