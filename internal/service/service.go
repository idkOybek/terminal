// internal/service/service.go

package service

import (
	"context"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, user *models.UserCreateRequest) (*models.User, error)
	GetByID(ctx context.Context, id int) (*models.User, error)
	Update(ctx context.Context, id int, user *models.UserUpdateRequest) (*models.User, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]models.User, error)
}

type FiscalModuleService interface {
	Create(ctx context.Context, module *models.FiscalModuleCreateRequest) (*models.FiscalModule, error)
	GetByID(ctx context.Context, id int) (*models.FiscalModule, error)
	Update(ctx context.Context, id int, module *models.FiscalModuleUpdateRequest) (*models.FiscalModule, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]models.FiscalModule, error)
}

type TerminalService interface {
	Create(ctx context.Context, terminal *models.TerminalCreateRequest) (*models.Terminal, error)
	GetByID(ctx context.Context, id int) (*models.Terminal, error)
	Update(ctx context.Context, id int, terminal *models.TerminalUpdateRequest) (*models.Terminal, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]models.Terminal, error)
}

type LinkService interface {
	Create(ctx context.Context, link *models.LinkCreateRequest) (*models.Link, error)
	GetByFactoryNumber(ctx context.Context, factoryNumber string) (*models.Link, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]models.Link, error)
}

type Services struct {
	UserService
	FiscalModuleService
	TerminalService
	LinkService
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		UserService:         NewUserService(deps.Repos.User),
		FiscalModuleService: NewFiscalModuleService(deps.Repos.FiscalModule),
		TerminalService:     NewTerminalService(deps.Repos.Terminal, deps.Repos.Link),
		LinkService:         NewLinkService(deps.Repos.Link),
	}
}