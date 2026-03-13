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

// UpdatePost updates the title and of an existing user.
// It returns an error if the update operation fails.
func (repository PostRepository) UpdatePost(post models.Post) error {
	stetment, err := repository.db.Prepare(
		`UPDATE posts SET title = $1, content = $2 WHERE id = $3`,
	)
	if err != nil {
		return err
	}
	defer stetment.Close()

	_, err = stetment.Exec(post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeletePost removes a post from the database based on its ID.
// It returns an error if the delete operation fails.
func (repository PostRepository) DeletePost(postID uint64) error {
	stetment, err := repository.db.Prepare(
		`DELETE FROM posts WHERE id = $1`,
	)
	if err != nil {
		return err
	}
	defer stetment.Close()

	_, err = stetment.Exec(postID)
	if err != nil {
		return err
	}

	return nil
}

// FindPostsByUser returns all posts created by a specific user.
// It retrieves the posts along with the author's username from the database
// based on the provided user ID.
func (repository PostRepository) FindPostsByUser(userID uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(
		`SELECT p.*, u.username from posts p
		join users u on u.id = p.author_id 
		where p.author_id = $1`,
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

func (repository PostRepository) LikePost(postID uint64) error {
	stetement, err := repository.db.Prepare(
		`UPDATE posts SET likes = likes + 1 WHERE id = $1`,
	)
	if err != nil {
		return err
	}
	defer stetement.Close()

	_, err = stetement.Exec(postID)
	if err != nil {
		return err
	}

	return nil
}
