package service

import (
	"clean-arch/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *models.Post) (*models.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*models.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]models.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]models.Post), args.Error(1)
}
func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	post := models.Post{Id: int64(1), Title: "Post 1", Text: "some text"}

	//Expect
	mockRepo.On("FindAll").Return([]models.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, int64(1), result[0].Id)
	assert.Equal(t, "Post 1", result[0].Title)
	assert.Equal(t, "some text", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	var id int64 = 1
	post := models.Post{Id: id, Title: "Post 1", Text: "some text"}

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, id, result.Id)
	assert.Equal(t, "Post 1", result.Title)
	assert.Equal(t, "some text", result.Text)
	assert.Nil(t, err)
}

func TestValidateEmptyPost(t *testing.T) {
	testservice := NewPostService(nil)
	err := testservice.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())
}

func TestValidateEmptyTitle(t *testing.T) {
	testservice := NewPostService(nil)
	post := models.Post{Id: 1, Title: "", Text: "some text"}
	err := testservice.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "Title can't be empty", err.Error())
}
