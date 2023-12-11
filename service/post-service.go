package service

import (
	"clean-arch/models"
	"clean-arch/repository"
	"errors"
	"math/rand"
)

type PostService interface {
	Validate(post *models.Post) error
	Create(post *models.Post) (*models.Post, error)
	FindAll() ([]models.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository = repository.NewPostgresRepository()
)

func NewPostService() PostService {
	return &service{}
}

func (*service) Validate(post *models.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("Title can't be empty")
		return err
	}
	return nil
}
func (*service) Create(post *models.Post) (*models.Post, error) {
	post.Id = rand.Int63()
	return repo.Save(post)
}
func (*service) FindAll() ([]models.Post, error) {
	return repo.FindAll()
}
