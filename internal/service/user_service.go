package service

import (
	"context"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo             repository.UserRepository
	fiscalModuleRepo repository.FiscalModuleRepository
}

func NewUserService(repo repository.UserRepository, fiscalModuleRepo repository.FiscalModuleRepository) *UserService {
	return &UserService{
		repo:             repo,
		fiscalModuleRepo: fiscalModuleRepo,
	}
}

func (s *UserService) Create(ctx context.Context, req *models.UserCreateRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		INN:      req.INN,
		Username: req.Username,
		Password: string(hashedPassword),
		IsActive: req.IsActive,
		IsAdmin:  req.IsAdmin,
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, id int, req *models.UserUpdateRequest) (*models.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.INN != nil {
		user.INN = *req.INN
	}
	if req.Username != nil {
		user.Username = *req.Username
	}
	if req.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}
	if req.IsAdmin != nil {
		user.IsAdmin = *req.IsAdmin
	}

	err = s.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	// Сначала удаляем связанные фискальные модули
	if err := s.fiscalModuleRepo.DeleteByUserID(ctx, id); err != nil {
		return err
	}

	// Затем удаляем самого пользователя
	return s.repo.Delete(ctx, id)
}

func (s *UserService) List(ctx context.Context) ([]*models.User, error) {
	return s.repo.List(ctx)
}
