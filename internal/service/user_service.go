// internal/service/user_service.go

package service

import (
	"context"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) Create(ctx context.Context, userReq *models.UserCreateRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		INN:      userReq.INN,
		Username: userReq.Username,
		Password: string(hashedPassword),
		IsActive: true,
		IsAdmin:  userReq.IsAdmin,
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) GetByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserServiceImpl) Update(ctx context.Context, id int, userReq *models.UserUpdateRequest) (*models.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if userReq.INN != "" {
		user.INN = userReq.INN
	}
	if userReq.Username != "" {
		user.Username = userReq.Username
	}
	if userReq.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}
	user.IsActive = userReq.IsActive
	user.IsAdmin = userReq.IsAdmin

	err = s.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserServiceImpl) List(ctx context.Context) ([]models.User, error) {
	return s.repo.List(ctx)
}
