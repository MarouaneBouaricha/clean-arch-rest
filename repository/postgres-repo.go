package repository

import (
	"clean-arch/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "tux"
	password = "mysecretpassword"
	dbname   = "blog"
)

type repo struct{}

func NewPostgresRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *models.Post) (bool, error) {
	db, err := connectToPostgreSQL()
	if err != nil {
		return false, err
	}
	// Perform database migration
	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		return false, err
	}

	err = db.Create(post).Error
	if err != nil {
		return false, err
	}
	log.Println("Created Post:", post)
	return true, nil
}

func (*repo) FindAll() ([]models.Post, error) {
	db, err := connectToPostgreSQL()
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	result := db.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func connectToPostgreSQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
