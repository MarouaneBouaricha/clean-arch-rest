package repository

import (
	"clean-arch/models"
)

type PostRepository interface {
	Save(post *models.Post) (bool, error)
	FindAll() ([]models.Post, error)
}
