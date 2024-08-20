package service

import (
    "github.com/idkOybek/newNewTerminal/internal/repository"
    "github.com/idkOybek/newNewTerminal/pkg/logger"
)

type Services struct {
    Auth         *AuthService
    User         *UserService
    FiscalModule *FiscalModuleService
    Terminal     *TerminalService
}

type Deps struct {
    Repos  *repository.Repositories
    Logger *logger.Logger
}

func NewServices(deps Deps) *Services {
    authService := NewAuthService(deps.Repos.User)
    userService := NewUserService(deps.Repos.User, deps.Repos.FiscalModule)
    fiscalModuleService := NewFiscalModuleService(deps.Repos.FiscalModule)
    terminalService := NewTerminalService(deps.Repos.Terminal, deps.Repos.FiscalModule, fiscalModuleService)

    return &Services{
        Auth:         authService,
        User:         userService,
        FiscalModule: fiscalModuleService,
        Terminal:     terminalService,
    }
}