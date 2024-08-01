package service

import (
	"context"
	"time"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type FiscalModuleService struct {
	repo repository.FiscalModuleRepository
}

func NewFiscalModuleService(repo repository.FiscalModuleRepository) *FiscalModuleService {
	return &FiscalModuleService{
		repo: repo,
	}
}

func (s *FiscalModuleService) Create(ctx context.Context, req *models.FiscalModuleCreateRequest) (*models.FiscalModule, error) {
	module := &models.FiscalModule{
		FiscalNumber:  req.FiscalNumber,
		FactoryNumber: req.FactoryNumber,
		UserID:        req.UserID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := s.repo.Create(ctx, module)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (s *FiscalModuleService) GetByID(ctx context.Context, id int) (*models.FiscalModule, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *FiscalModuleService) GetByFactoryNumber(ctx context.Context, factoryNumber string) (*models.FiscalModule, error) {
	return s.repo.GetByFactoryNumber(ctx, factoryNumber)
}

func (s *FiscalModuleService) Update(ctx context.Context, id int, req *models.FiscalModuleUpdateRequest) (*models.FiscalModule, error) {
	module, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	module.FiscalNumber = req.FiscalNumber
	module.FactoryNumber = req.FactoryNumber
	module.UserID = req.UserID
	module.UpdatedAt = time.Now()

	err = s.repo.Update(ctx, module)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (s *FiscalModuleService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *FiscalModuleService) List(ctx context.Context) ([]*models.FiscalModule, error) {
	return s.repo.List(ctx)
}
