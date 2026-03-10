package models

import (
	"errors"
	"strings"
	"time"
)

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

// Prepare validates and formats the post data before saving.
func (p *Post) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}

	p.format()
	return nil
}

func (p Post) validate() error {
	if p.Title == "" {
		return errors.New("Title is required.")
	}

	if p.Content == "" {
		return errors.New("Content is required.")
	}

	if len(p.Title) < 5 {
		return errors.New("Title require min 5 caracteres.")
	}

	if len(p.Content) < 10 {
		return errors.New("Content require min 10 caracteres.")
	}

	return nil
}

func (p *Post) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}
