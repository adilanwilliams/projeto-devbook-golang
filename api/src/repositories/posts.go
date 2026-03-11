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

// FindUserFeed returns the posts created by the user and by the users they follow.
// It retrieves the feed ordered by the most recent posts.
func (repository PostRepository) FindUserFeed(userID uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(
		`SELECT DISTINCT p.*, u.username FROM posts p 
		INNER JOIN users u ON u.id = p.author_id 
		INNER JOIN follows f ON p.author_id = f.user_id 
		WHERE u.id = $1 or f.follow_id = $1
		ORDER BY p.created_at`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

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
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
