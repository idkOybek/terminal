package service

import (
	"context"
	"errors"
	"time"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type TerminalService struct {
	repo             repository.TerminalRepository
	fiscalModuleRepo repository.FiscalModuleRepository
}

func NewTerminalService(repo repository.TerminalRepository, fiscalModuleRepo repository.FiscalModuleRepository) *TerminalService {
	return &TerminalService{
		repo:             repo,
		fiscalModuleRepo: fiscalModuleRepo,
	}
}

func (s *TerminalService) Create(ctx context.Context, req *models.TerminalCreateRequest) (*models.Terminal, error) {
	// Check if fiscal module exists for the given assembly number
	_, err := s.fiscalModuleRepo.GetByFactoryNumber(ctx, req.AssemblyNumber)
	if err != nil {
		return nil, errors.New("no fiscal module found for the given assembly number")
	}

	terminal := &models.Terminal{
		AssemblyNumber:     req.AssemblyNumber,
		INN:                req.INN,
		CompanyName:        req.CompanyName,
		Address:            req.Address,
		CashRegisterNumber: req.CashRegisterNumber,
		Status:             true,
		UserID:             req.UserID,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	err = s.repo.Create(ctx, terminal)
	if err != nil {
		return nil, err
	}

	return terminal, nil
}

func (s *TerminalService) GetByID(ctx context.Context, id int) (*models.Terminal, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TerminalService) Update(ctx context.Context, id int, req *models.TerminalUpdateRequest) (*models.Terminal, error) {
	terminal, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	terminal.INN = req.INN
	terminal.CompanyName = req.CompanyName
	terminal.Address = req.Address
	terminal.CashRegisterNumber = req.CashRegisterNumber
	terminal.Status = req.Status
	terminal.UpdatedAt = time.Now()

	err = s.repo.Update(ctx, terminal)
	if err != nil {
		return nil, err
	}

	return terminal, nil
}

func (s *TerminalService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *TerminalService) List(ctx context.Context) ([]*models.Terminal, error) {
	return s.repo.List(ctx)
}
