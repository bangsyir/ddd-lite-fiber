package user

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name, email string) (*User, error) {
	user := &User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) FindByID(ctx context.Context, id string) (*User, error) {
	return s.repo.FindById(ctx, id)
}

func (s *Service) Update(ctx context.Context, id, name, email string) (*User, error) {
	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	user.Name = name
	user.Email = email
	user.UpdatedAt = time.Now()

	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*User, error) {
	return s.repo.FindAll(ctx)
}
