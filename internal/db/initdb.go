package db

import (
	"os"

	"github.com/glebpepega/new1/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func New() *DB {
	return &DB{}
}

func (db *DB) Init() error {
	dbInstance, err := gorm.Open(postgres.Open(os.Getenv("DB_CONN_INFO")), &gorm.Config{})
	if err != nil {
		return err
	}
	db.DB = dbInstance
	if err := db.constructTables(); err != nil {
		return err
	}
	return nil
}

func (db *DB) constructTables() error {
	if err := db.DB.AutoMigrate(&models.News{}); err != nil {
		return err
	}
	if err := db.DB.AutoMigrate(&models.NewsCategories{}); err != nil {
		return err
	}
	return nil
}
