// internal/service/fiscal_module_service.go

package service

import (
	"context"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type FiscalModuleServiceImpl struct {
	repo repository.FiscalModuleRepository
}

func NewFiscalModuleService(repo repository.FiscalModuleRepository) *FiscalModuleServiceImpl {
	return &FiscalModuleServiceImpl{
		repo: repo,
	}
}

func (s *FiscalModuleServiceImpl) Create(ctx context.Context, moduleReq *models.FiscalModuleCreateRequest) (*models.FiscalModule, error) {
	module := &models.FiscalModule{
		FactoryNumber: moduleReq.FactoryNumber,
		FiscalNumber:  moduleReq.FiscalNumber,
		UserID:        moduleReq.UserID,
	}

	err := s.repo.Create(ctx, module)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (s *FiscalModuleServiceImpl) GetByID(ctx context.Context, id int) (*models.FiscalModule, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *FiscalModuleServiceImpl) Update(ctx context.Context, id int, moduleReq *models.FiscalModuleUpdateRequest) (*models.FiscalModule, error) {
	module, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if moduleReq.FactoryNumber != "" {
		module.FactoryNumber = moduleReq.FactoryNumber
	}
	if moduleReq.FiscalNumber != "" {
		module.FiscalNumber = moduleReq.FiscalNumber
	}
	if moduleReq.UserID != 0 {
		module.UserID = moduleReq.UserID
	}

	err = s.repo.Update(ctx, module)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (s *FiscalModuleServiceImpl) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *FiscalModuleServiceImpl) List(ctx context.Context) ([]models.FiscalModule, error) {
	return s.repo.List(ctx)
}
