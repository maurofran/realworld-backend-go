package user

import (
	"context"
)

type Model struct {
	ID        interface{} `bson:"_id,omitempty"`
	Username  string      `bson:"username"`
	Password  string      `bson:"password"`
	Email     string      `bson:"email"`
	Biography string      `bson:"biography"`
	Image     string      `bson:"image"`
}

type Repository interface {
	Add(context.Context, *Model) error
	Update(context.Context, *Model) error
	GetByUsername(context.Context, string) (*Model, error)
}
