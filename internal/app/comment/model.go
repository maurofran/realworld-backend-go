package comment

import "time"

type Comment struct {
	ID        interface{} `bson:"_id,omitempty"`
	Version   int         `bson:"_v"`
	ArticleID interface{} `bson:"articleId"`
	Body      string      `bson:"body"`
	AuthorID  interface{} `bson:"authorId"`
	CreatedAt time.Time   `bson:"createdAt"`
	UpdatedAt time.Time   `bson:"updatedAt"`
}

type Repository interface {
	Add(*Comment) error
	Update(*Comment) error
	Remove(*Comment) error
	GetByID(interface{}) (*Comment, error)
	GetByArticle(interface{}) ([]*Comment, error)
}
