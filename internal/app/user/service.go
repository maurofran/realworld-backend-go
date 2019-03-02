package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"realworld/internal/app/security"
)

type Service interface {
	Login(context.Context, *LoginCommand) (*Representation, error)
	Register(context.Context, *RegisterCommand) (*Representation, error)
	Update(context.Context, *UpdateCommand) (*Representation, error)
	GetCurrent(context.Context) (*Representation, error)
	GetByUsername(ctx context.Context, username string) (*Representation, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Login(ctx context.Context, cmd *LoginCommand) (*Representation, error) {
	user, err := s.repository.GetByUsername(ctx, cmd.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, notFound(cmd.Username)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cmd.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, wrongPassword(cmd.Username)
	}
	if err != nil {
		return nil, err
	}
	token, err := security.CreateToken(user.Username, user.Email)
	if err != nil {
		return nil, err
	}
	representation := representationOf(user)
	representation.Token = token
	return representation, nil
}

func (s *service) Register(ctx context.Context, cmd *RegisterCommand) (*Representation, error) {
	user, err := s.repository.GetByUsername(ctx, cmd.Username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, alreadyExists(cmd.Username)
	}
	password, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user = &Model{
		Username: cmd.Username,
		Password: string(password),
		Email:    cmd.Email,
	}
	if err := s.repository.Add(ctx, user); err != nil {
		return nil, err
	}
	return representationOf(user), nil
}

func (s *service) Update(ctx context.Context, cmd *UpdateCommand) (*Representation, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	user.Email = cmd.Email
	user.Biography = cmd.Biography
	user.Image = cmd.Image
	if err := s.repository.Update(ctx, user); err != nil {
		return nil, err
	}
	return representationOf(user), nil
}

func (s *service) GetCurrent(ctx context.Context) (*Representation, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	return representationOf(user), nil
}

func (s *service) GetByUsername(ctx context.Context, username string) (*Representation, error) {
	user, err := s.repository.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, notFound(username)
	}
	return representationOf(user), nil
}

func (s *service) currentUser(ctx context.Context) (*Model, error) {
	username, err := security.Username(ctx)
	if err != nil {
		return nil, err
	}
	user, err := s.repository.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, notFound(username)
	}
	return user, nil
}

func representationOf(user *Model) *Representation {
	return &Representation{
		Username:  user.Username,
		Email:     user.Email,
		Biography: user.Biography,
		Image:     user.Image,
	}
}
