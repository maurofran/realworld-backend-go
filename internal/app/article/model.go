package article

import (
	"context"
	"github.com/gosimple/slug"
	"time"
)

type Article struct {
	ID          interface{} `bson:"_id,omitempty"`
	Version     int         `bson:"_v"`
	Slug        string      `bson:"slug"`
	Title       string      `bson:"title"`
	Description string      `bson:"title"`
	Body        string      `bson:"body"`
	Tags        []string    `bson:"tags"`
	Author      string      `bson:"author"`
	CreatedAt   time.Time   `bson:"createdAt"`
	UpdatedAt   time.Time   `bson:"updatedAt"`
}

func New(title, description, body string, tags []string, author string) (*Article, error) {
	// TODO Validate parameters.
	return &Article{
		Slug:   slug.Make(title),
		Title:  title,
		Body:   body,
		Tags:   tags,
		Author: author,
	}, nil
}

type Repository interface {
	Add(context.Context, *Article) error
	Update(context.Context, *Article) error
	Remove(context.Context, *Article) error
	GetBySlug(context.Context, string) (*Article, error)
	GetAll(ctx context.Context, tag, author, favorited string, offset, size int) ([]*Article, int, error)
}
