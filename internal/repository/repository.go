package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository/postgres"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
)

type Repositories struct {
	User         UserRepository
	FiscalModule FiscalModuleRepository
	Terminal     TerminalRepository
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*models.User, error)
}

type FiscalModuleRepository interface {
	Create(ctx context.Context, module *models.FiscalModule) error
	GetByID(ctx context.Context, id int) (*models.FiscalModule, error)
	GetByFactoryNumber(ctx context.Context, factoryNumber string) (*models.FiscalModule, error)
	Update(ctx context.Context, module *models.FiscalModule) error
	Delete(ctx context.Context, id int) error
	DeleteByUserID(ctx context.Context, userID int) error
	List(ctx context.Context) ([]*models.FiscalModule, error)
}

type TerminalRepository interface {
	Create(ctx context.Context, terminal *models.Terminal) error
	GetByID(ctx context.Context, id int) (*models.Terminal, error)
	Update(ctx context.Context, terminal *models.Terminal) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*models.Terminal, error)
	GetUserIDByCashRegisterNumber(ctx context.Context, cashRegisterNumber string) (int, error)
}

func NewRepositories(db *sql.DB, logger *logger.Logger) *Repositories {
	if db == nil {
		log.Fatal("Database connection is nil")
	}
	if logger == nil {
		log.Fatal("Logger is nil")
	}
	return &Repositories{
		User:         postgres.NewUserRepository(db, logger),
		FiscalModule: postgres.NewFiscalModuleRepository(db, logger),
		Terminal:     postgres.NewTerminalRepository(db, logger),
	}
}

// Интерфейсы для функций создания репозиториев
type UserRepoCreator func(*sql.DB, *logger.Logger) UserRepository
type FiscalModuleRepoCreator func(*sql.DB, *logger.Logger) FiscalModuleRepository
type TerminalRepoCreator func(*sql.DB, *logger.Logger) TerminalRepository
