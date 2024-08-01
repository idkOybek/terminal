package service

import (
	"context"
	"time"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user *models.UserCreateRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Username:  user.Username,
		Password:  string(hashedPassword),
		IsAdmin:   user.IsAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *UserService) GetByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, user *models.UserUpdateRequest) (*models.User, error) {
	existingUser, err := s.repo.GetByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	if user.Username != "" {
		existingUser.Username = user.Username
	}

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		existingUser.Password = string(hashedPassword)
	}

	existingUser.IsAdmin = user.IsAdmin
	existingUser.UpdatedAt = time.Now()

	err = s.repo.Update(ctx, existingUser)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) List(ctx context.Context) ([]*models.User, error) {
	return s.repo.List(ctx)
}
