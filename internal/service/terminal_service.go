package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type TerminalService struct {
	repo                repository.TerminalRepository
	fiscalModuleRepo    repository.FiscalModuleRepository
	fiscalModuleService *FiscalModuleService // Добавляем сервис фискального модуля
}

func NewTerminalService(repo repository.TerminalRepository, fiscalModuleRepo repository.FiscalModuleRepository, fiscalModuleService *FiscalModuleService) *TerminalService {
	return &TerminalService{
		repo:                repo,
		fiscalModuleRepo:    fiscalModuleRepo,
		fiscalModuleService: fiscalModuleService,
	}
}

func (s *TerminalService) Create(ctx context.Context, req *models.TerminalCreateRequest) (*models.Terminal, error) {
	fiscalModule, err := s.fiscalModuleRepo.GetByFactoryNumber(ctx, req.CashRegisterNumber)
	if err != nil {
		return nil, err
	}
	if fiscalModule == nil {
		return nil, errors.New("no fiscal module found with the given cash register number")
	}
	userID, err := s.repo.GetUserIDByCashRegisterNumber(ctx, req.CashRegisterNumber)
	if err != nil {
		return nil, errors.New("failed to determine user for this terminal")
	}

	lastRequestDate, _ := time.Parse(time.RFC3339, req.LastRequestDate)
	databaseUpdateDate, _ := time.Parse(time.RFC3339, req.DatabaseUpdateDate)

	terminal := &models.Terminal{
		AssemblyNumber:     req.AssemblyNumber,
		INN:                req.INN,
		CompanyName:        req.CompanyName,
		Address:            req.Address,
		CashRegisterNumber: req.CashRegisterNumber,
		ModuleNumber:       req.ModuleNumber,
		LastRequestDate:    lastRequestDate,
		DatabaseUpdateDate: databaseUpdateDate,
		IsActive:           true,
		UserID:             userID,
		FreeRecordBalance:  req.FreeRecordBalance,
	}

	err = s.repo.Create(ctx, terminal)
	if err != nil {
		return nil, err
	}

	// Активируем связанный фискальный модуль
	err = s.fiscalModuleService.Activate(ctx, fiscalModule.ID)
	if err != nil {
		return nil, errors.New(fmt.Sprint("failed to activate fiscal module: %w", err))
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

	// Обновляем только те поля, которые присутствуют в запросе
	if req.AssemblyNumber != nil {
		terminal.AssemblyNumber = *req.AssemblyNumber
	}
	if req.INN != nil {
		terminal.INN = *req.INN
	}
	if req.CompanyName != nil {
		terminal.CompanyName = *req.CompanyName
	}
	if req.Address != nil {
		terminal.Address = *req.Address
	}
	if req.CashRegisterNumber != nil {
		terminal.CashRegisterNumber = *req.CashRegisterNumber
	}
	if req.ModuleNumber != nil {
		terminal.ModuleNumber = *req.ModuleNumber
	}
	if req.LastRequestDate != nil {
		lastRequestDate, err := time.Parse(time.RFC3339, *req.LastRequestDate)
		if err == nil {
			terminal.LastRequestDate = lastRequestDate
		}
	}
	if req.DatabaseUpdateDate != nil {
		databaseUpdateDate, err := time.Parse(time.RFC3339, *req.DatabaseUpdateDate)
		if err == nil {
			terminal.DatabaseUpdateDate = databaseUpdateDate
		}
	}
	if req.IsActive != nil {
		terminal.IsActive = *req.IsActive
	}
	if req.UserID != nil {
		terminal.UserID = *req.UserID
	}
	if req.FreeRecordBalance != nil {
		terminal.FreeRecordBalance = *req.FreeRecordBalance
	}

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
