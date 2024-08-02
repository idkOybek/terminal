package service

import (
	"context"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
	"github.com/idkOybek/newNewTerminal/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(ctx context.Context, req *models.UserCreateRequest) (*models.User, error) {
	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		INN:      req.INN,
		IsActive: req.IsActive,
		IsAdmin:  req.IsAdmin,
	}

	// Сохраняем пользователя в базу данных
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	// Не возвращаем хешированный пароль
	user.Password = ""

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, req *models.UserLoginRequest) (*models.UserLoginResponse, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	// Проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	// Генерируем JWT токен
	token, err := auth.GenerateToken(user.ID, user.Username, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	return &models.UserLoginResponse{
		User:  *user,
		Token: token,
	}, nil
}
