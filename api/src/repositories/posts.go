package repositories

import (
	"database/sql"
	"devbook/src/models"
)

// PostRepository provides access to post persistence operations.
type PostRepository struct {
	db *sql.DB
}

// NewPostRepository creates and returns a new instance of PostRepository.
func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db}
}

// Save inserts a new post into the database and returns the generated ID.
// It returns an error if the insert operation fails.
func (repository PostRepository) Save(post models.Post, userID uint64) (uint64, error) {
	var id uint64

	err := repository.db.QueryRow(
		`INSERT INTO posts 
		(title, author_id, content, likes)
		values ($1, $2, $3, $4) RETURNING id`,
		post.Title,
		userID,
		post.Content,
		post.Likes,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	repository.db.Close()

	return id, nil
}
