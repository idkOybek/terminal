// internal/repository/repository.go

package repository

import (
	"context"

	"github.com/idkOybek/newNewTerminal/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]models.User, error)
}

type FiscalModuleRepository interface {
	Create(ctx context.Context, module *models.FiscalModule) error
	GetByID(ctx context.Context, id int) (*models.FiscalModule, error)
	GetByFiscalNumber(ctx context.Context, fiscalNumber string) (*models.FiscalModule, error)
	Update(ctx context.Context, module *models.FiscalModule) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]models.FiscalModule, error)
}

type TerminalRepository interface {
	Create(ctx context.Context, terminal *models.Terminal) error
	GetByID(ctx context.Context, id int) (*models.Terminal, error)
	Update(ctx context.Context, terminal *models.Terminal) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]models.Terminal, error)
}

type LinkRepository interface {
	Create(ctx context.Context, link *models.Link) error
	GetByFactoryNumber(ctx context.Context, factoryNumber string) (*models.Link, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]models.Link, error)
}