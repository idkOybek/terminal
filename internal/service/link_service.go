// internal/service/link_service.go

package service

import (
	"context"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type LinkServiceImpl struct {
	repo repository.LinkRepository
}

func NewLinkService(repo repository.LinkRepository) *LinkServiceImpl {
	return &LinkServiceImpl{
		repo: repo,
	}
}

func (s *LinkServiceImpl) Create(ctx context.Context, linkReq *models.LinkCreateRequest) (*models.Link, error) {
	link := &models.Link{
		FiscalNumber:  linkReq.FiscalNumber,
		FactoryNumber: linkReq.FactoryNumber,
	}

	err := s.repo.Create(ctx, link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (s *LinkServiceImpl) GetByFactoryNumber(ctx context.Context, factoryNumber string) (*models.Link, error) {
	return s.repo.GetByFactoryNumber(ctx, factoryNumber)
}

func (s *LinkServiceImpl) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *LinkServiceImpl) List(ctx context.Context) ([]models.Link, error) {
	return s.repo.List(ctx)
}