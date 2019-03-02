package comment

import (
	"realworld/internal/app/article"
	"time"
)

type Collection struct {
	comments []*Representation `json:"comments"`
}

type Representation struct {
	ID        interface{}     `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Body      string          `json:"body"`
	Author    *article.Author `json:"author"`
}
