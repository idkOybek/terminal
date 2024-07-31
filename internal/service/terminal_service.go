// internal/service/terminal_service.go

package service

import (
	"context"
	"errors"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type TerminalServiceImpl struct {
	repo     repository.TerminalRepository
	linkRepo repository.LinkRepository
}

func NewTerminalService(repo repository.TerminalRepository, linkRepo repository.LinkRepository) *TerminalServiceImpl {
	return &TerminalServiceImpl{
		repo:     repo,
		linkRepo: linkRepo,
	}
}

func (s *TerminalServiceImpl) Create(ctx context.Context, terminalReq *models.TerminalCreateRequest) (*models.Terminal, error) {
	// Проверяем наличие связки
	link, err := s.linkRepo.GetByFactoryNumber(ctx, terminalReq.ModuleNumber)
	if err != nil {
		return nil, errors.New("invalid module number: no link found")
	}

	terminal := &models.Terminal{
		INN:                terminalReq.INN,
		CompanyName:        terminalReq.CompanyName,
		Address:            terminalReq.Address,
		CashRegisterNumber: terminalReq.CashRegisterNumber,
		ModuleNumber:       terminalReq.ModuleNumber,
		AssemblyNumber:     terminalReq.AssemblyNumber,
		Status:             true,
		UserID:             terminalReq.UserID,
		FreeRecordBalance:  terminalReq.FreeRecordBalance,
	}

	err = s.repo.Create(ctx, terminal)
	if err != nil {
		return nil, err
	}

	return terminal, nil
}

func (s *TerminalServiceImpl) GetByID(ctx context.Context, id int) (*models.Terminal, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TerminalServiceImpl) Update(ctx context.Context, id int, terminalReq *models.TerminalUpdateRequest) (*models.Terminal, error) {
	terminal, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if terminalReq.INN != "" {
		terminal.INN = terminalReq.INN
	}
	if terminalReq.CompanyName != "" {
		terminal.CompanyName = terminalReq.CompanyName
	}
	if terminalReq.Address != "" {
		terminal.Address = terminalReq.Address
	}
	if terminalReq.CashRegisterNumber != "" {
		terminal.CashRegisterNumber = terminalReq.CashRegisterNumber
	}
	if terminalReq.ModuleNumber != "" {
		// Проверяем наличие связки при изменении номера модуля
		_, err := s.linkRepo.GetByFactoryNumber(ctx, terminalReq.ModuleNumber)
		if err != nil {
			return nil, errors.New("invalid module number: no link found")
		}
		terminal.ModuleNumber = terminalReq.ModuleNumber
	}
	if terminalReq.AssemblyNumber != "" {
		terminal.AssemblyNumber = terminalReq.AssemblyNumber
	}
	terminal.Status = terminalReq.Status
	if terminalReq.UserID != 0 {
		terminal.UserID = terminalReq.UserID
	}
	if terminalReq.FreeRecordBalance != 0 {
		terminal.FreeRecordBalance = terminalReq.FreeRecordBalance
	}

	err = s.repo.Update(ctx, terminal)
	if err != nil {
		return nil, err
	}

	return terminal, nil
}

func (s *TerminalServiceImpl) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *TerminalServiceImpl) List(ctx context.Context) ([]models.Terminal, error) {
	return s.repo.List(ctx)
}
