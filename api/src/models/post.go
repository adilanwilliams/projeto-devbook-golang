package models

import "time"

// Post represents a post entity created by a user and stored in the database.
type Post struct {
	ID        uint64    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	AuthorID  uint64    `json:"author_id,omitempty"`
	Author    string    `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Likes     uint64    `json:"likes"`
}