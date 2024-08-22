package service

import (
	"context"
	"fmt"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
)

type FiscalModuleService struct {
	repo   repository.FiscalModuleRepository
	logger *logger.Logger
}

func NewFiscalModuleService(repo repository.FiscalModuleRepository, logger *logger.Logger) *FiscalModuleService {
	return &FiscalModuleService{
		repo:   repo,
		logger: logger,
	}
}
func (s *FiscalModuleService) Create(ctx context.Context, req *models.FiscalModuleCreateRequest) (*models.FiscalModuleResponse, error) {
	module := &models.FiscalModule{
		FiscalNumber:  req.FiscalNumber,
		FactoryNumber: req.FactoryNumber,
		UserID:        req.UserID,
		IsActive:      req.IsActive, // Новое поле
	}

	err := s.repo.Create(ctx, module)
	if err != nil {
		return nil, err
	}

	return &models.FiscalModuleResponse{
		ID:            module.ID,
		FiscalNumber:  module.FiscalNumber,
		FactoryNumber: module.FactoryNumber,
		UserID:        module.UserID,
		IsActive:      module.IsActive, // Новое поле
	}, nil
}

func (s *FiscalModuleService) GetByID(ctx context.Context, id int) (*models.FiscalModuleResponse, error) {
	module, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.FiscalModuleResponse{
		ID:            module.ID,
		FiscalNumber:  module.FiscalNumber,
		FactoryNumber: module.FactoryNumber,
		UserID:        module.UserID,
	}, nil
}

func (s *FiscalModuleService) Update(ctx context.Context, id int, req *models.FiscalModuleUpdateRequest) (*models.FiscalModule, error) {
	module, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.FiscalNumber != nil {
		module.FiscalNumber = *req.FiscalNumber
	}
	if req.FactoryNumber != nil {
		module.FactoryNumber = *req.FactoryNumber
	}
	if req.UserID != nil {
		module.UserID = *req.UserID
	}
	if req.IsActive != nil {
		module.IsActive = *req.IsActive
	}

	err = s.repo.Update(ctx, module)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (s *FiscalModuleService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *FiscalModuleService) List(ctx context.Context) ([]*models.FiscalModuleResponse, error) {
	modules, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var response []*models.FiscalModuleResponse
	for _, module := range modules {
		response = append(response, &models.FiscalModuleResponse{
			ID:            module.ID,
			FiscalNumber:  module.FiscalNumber,
			FactoryNumber: module.FactoryNumber,
			UserID:        module.UserID,
			IsActive:      module.IsActive,
		})
	}

	return response, nil
}

func (s *FiscalModuleService) Activate(ctx context.Context, id int) error {
	s.logger.Info("Starting fiscal module activation", "id", id)

	module, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get fiscal module", "id", id, "error", err)
		return fmt.Errorf("failed to get fiscal module: %w", err)
	}
	s.logger.Info("Fiscal module retrieved", "id", id, "current_status", module.IsActive)

	if !module.IsActive {
		module.IsActive = true
		s.logger.Info("Updating fiscal module status", "id", id, "new_status", true)
		err = s.repo.Update(ctx, module)
		if err != nil {
			s.logger.Error("Failed to update fiscal module", "id", id, "error", err)
			return fmt.Errorf("failed to update fiscal module: %w", err)
		}
		s.logger.Info("Fiscal module activated successfully", "id", id)
	} else {
		s.logger.Info("Fiscal module already active", "id", id)
	}

	return nil
}
