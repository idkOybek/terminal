package service

import (
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type Services struct {
	Auth         *AuthService
	User         *UserService
	FiscalModule *FiscalModuleService
	Terminal     *TerminalService
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		Auth:         NewAuthService(deps.Repos.User),
		User:         NewUserService(deps.Repos.User),
		FiscalModule: NewFiscalModuleService(deps.Repos.FiscalModule),
		Terminal:     NewTerminalService(deps.Repos.Terminal),
	}
}
