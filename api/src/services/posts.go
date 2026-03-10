package services

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
)

// PostService provides business logic related to posts.
type PostService struct {
	PostRepository *repositories.PostRepository
}

// NewPostService creates and returns a new instance of PostService.
// It establishes a database connection and injects the PostRepository dependency.
func NewPostService() (*PostService, error) {
	db, err := database.Connect()
	if err != nil {
		return &PostService{}, err
	}

	repository := repositories.NewPostRepository(db)
	return &PostService{repository}, nil
}

// SavePost creates a new post in the database and returns the generated ID.
// It returns an error if the operation fails.
func (service PostService) SavePost(post models.Post,) (uint64, error) {
	err := post.Prepare()
	if err != nil {
		return 0, err
	}

	return service.PostRepository.Save(post)
}

func (service PostService) FindPostByID(postID uint64) (models.Post, error) {
	return service.PostRepository.FindPostByID(postID)
}