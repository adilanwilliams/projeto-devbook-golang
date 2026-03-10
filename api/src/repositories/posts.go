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
func (repository PostRepository) Save(post models.Post) (uint64, error) {
	var id uint64

	err := repository.db.QueryRow(
		`INSERT INTO posts 
		(title, author_id, content, likes)
		values ($1, $2, $3, $4) RETURNING id`,
		post.Title,
		post.AuthorID,
		post.Content,
		post.Likes,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	repository.db.Close()

	return id, nil
}

func (repository PostRepository) FindPostByID(postID uint64) (models.Post, error) {
	rows, err := repository.db.Query(
		`SELECT p.*, u.username 
		FROM posts p INNER JOIN users u 
		ON u.id = p.author_id WHERE p.id = $1`,
		postID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer rows.Close()

	var post models.Post
	if rows.Next() {
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.AuthorID,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.Author,
		)
		if err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

// func (repository PostRepository) FindPostByID(postID uint64) (models.Post, error) {
// 	rows, err := repository.db.Query(
// 		`SELECT 
// 		title, content, author_id, likes, created_at 
// 		FROM posts WHERE id = $1`,
// 		postID,
// 	)
// 	if err != nil {
// 		return models.Post{}, err
// 	}
// 	defer rows.Close()

// 	var post models.Post
// 	if rows.Next() {
// 		err = rows.Scan(
// 			&post.Title,
// 			&post.Content,
// 			&post.AuthorID,
// 			&post.Likes,
// 			&post.CreatedAt,
// 		)
// 		if err != nil {
// 			return models.Post{}, err
// 		}
// 	}

// 	return post, nil
// }
