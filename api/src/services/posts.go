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
func (service PostService) SavePost(post models.Post) (uint64, error) {
	err := post.Prepare()
	if err != nil {
		return 0, err
	}

	return service.PostRepository.Save(post)
}

// FindPostByID retrieves a post from the repository using its unique ID.
func (service PostService) FindPostByID(postID uint64) (models.Post, error) {
	return service.PostRepository.FindPostByID(postID)
}

// FindUserFeed retrieves the feed for a specific user.
// The feed includes posts created by the user and by the users they follow.
func (service PostService) FindUserFeed(userID uint64) ([]models.Post, error) {
	return service.PostRepository.FindUserFeed(userID)
}

// UpdatePost updates an existing post's information.
func (service PostService) UpdatePost(post models.Post) error {
	return service.PostRepository.UpdatePost(post)
}

// DeletePost removes a post from the database by its ID.
func (service PostService) DeletePost(postID uint64) error {
	return service.PostRepository.DeletePost(postID)
}

// FindPostsByUser returns all posts created by a specific user.
// It retrieves the posts along with the author's username from the database
// based on the provided user ID.
func (service PostService) FindPostsByUser(userID uint64) ([]models.Post, error) {
	return service.PostRepository.FindPostsByUser(userID)
}
