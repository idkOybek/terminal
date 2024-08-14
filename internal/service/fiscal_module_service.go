package service

import (
	"context"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type FiscalModuleService struct {
	repo repository.FiscalModuleRepository
}

func NewFiscalModuleService(repo repository.FiscalModuleRepository) *FiscalModuleService {
	return &FiscalModuleService{repo: repo}
}

func (s *FiscalModuleService) Create(ctx context.Context, req *models.FiscalModuleCreateRequest) (*models.FiscalModuleResponse, error) {
	module := &models.FiscalModule{
		FiscalNumber:  req.FiscalNumber,
		FactoryNumber: req.FactoryNumber,
		UserID:        req.UserID,
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
		})
	}

	return response, nil
}
