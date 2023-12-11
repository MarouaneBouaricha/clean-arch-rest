package repository

import (
	"clean-arch/models"
)

type PostRepository interface {
	Save(post *models.Post) (*models.Post, error)
	FindAll() ([]models.Post, error)
}
