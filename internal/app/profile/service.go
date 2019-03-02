package profile

import (
	"context"
	"realworld/internal/app/follower"
	"realworld/internal/app/user"
)

type Service interface {
	GetByUsername(ctx context.Context, username string) (*Representation, error)
	Follow(ctx context.Context, username string) (*Representation, error)
	Unfollow(ctx context.Context, username string) (*Representation, error)
}

type service struct {
	userService  user.Service
	followerRepo follower.Repository
}

func NewService(userService user.Service, followerRepo follower.Repository) Service {
	return &service{
		userService:  userService,
		followerRepo: followerRepo,
	}
}

func (s *service) GetByUsername(ctx context.Context, username string) (*Representation, error) {
	me, other, err := s.meAndOther(ctx, username)
	if err != nil {
		return nil, err
	}
	following, err := s.followerRepo.GetOne(ctx, me.Username, other.Username)
	if err != nil {
		return nil, err
	}
	return representationOf(other, following != nil), nil
}

func (s *service) Follow(ctx context.Context, username string) (*Representation, error) {
	me, other, err := s.meAndOther(ctx, username)
	if err != nil {
		return nil, err
	}
	following, err := s.followerRepo.GetOne(ctx, me.Username, other.Username)
	if err != nil {
		return nil, err
	}
	if following == nil {
		following, err = follower.New(me.Username, other.Username)
		if err != nil {
			return nil, err
		}
		if err := s.followerRepo.Add(ctx, following); err != nil {
			return nil, err
		}
	}
	return representationOf(other, true), nil
}

func (s *service) Unfollow(ctx context.Context, username string) (*Representation, error) {
	me, other, err := s.meAndOther(ctx, username)
	if err != nil {
		return nil, err
	}
	following, err := s.followerRepo.GetOne(ctx, me.Username, other.Username)
	if err != nil {
		return nil, err
	}
	if following != nil {
		if err := s.followerRepo.Remove(ctx, following); err != nil {
			return nil, err
		}
	}
	return representationOf(other, false), nil
}

func (s *service) meAndOther(ctx context.Context, username string) (*user.Representation, *user.Representation, error) {
	me, err := s.userService.GetCurrent(ctx)
	if err != nil {
		return nil, nil, err
	}
	other, err := s.userService.GetByUsername(ctx, username)
	if err != nil {
		return nil, nil, err
	}
	return me, other, nil
}

func representationOf(user *user.Representation, following bool) *Representation {
	return &Representation{
		Username:  user.Username,
		Image:     user.Image,
		Biography: user.Biography,
		Following: following,
	}
}
