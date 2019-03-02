package follower

import "context"

type Model struct {
	ID       interface{} `bson:"_id,omitempty"`
	Follower string      `bson:"follower"`
	Followee string      `bson:"followee"`
}

func New(follower, followee string) (*Model, error) {
	// TODO Validate arguments
	return &Model{
		Follower: follower,
		Followee: followee,
	}, nil
}

type Repository interface {
	Add(context.Context, *Model) error
	Remove(context.Context, *Model) error
	GetOne(ctx context.Context, follower, followee string) (*Model, error)
}
