package service

import (
    "context"
    "time"
    "github.com/idkOybek/newNewTerminal/internal/models"
    "github.com/idkOybek/newNewTerminal/internal/repository"
)

type TerminalService struct {
    repo repository.TerminalRepository
}

func NewTerminalService(repo repository.TerminalRepository) *TerminalService {
    return &TerminalService{repo: repo}
}

func (s *TerminalService) Create(ctx context.Context, req *models.TerminalCreateRequest) (*models.Terminal, error) {
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
        Status:             req.Status,
        UserID:             req.UserID,
        FreeRecordBalance:  req.FreeRecordBalance,
    }

    err := s.repo.Create(ctx, terminal)
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

    lastRequestDate, _ := time.Parse(time.RFC3339, req.LastRequestDate)
    databaseUpdateDate, _ := time.Parse(time.RFC3339, req.DatabaseUpdateDate)

    terminal.AssemblyNumber = req.AssemblyNumber
    terminal.INN = req.INN
    terminal.CompanyName = req.CompanyName
    terminal.Address = req.Address
    terminal.CashRegisterNumber = req.CashRegisterNumber
    terminal.ModuleNumber = req.ModuleNumber
    terminal.LastRequestDate = lastRequestDate
    terminal.DatabaseUpdateDate = databaseUpdateDate
    terminal.Status = req.Status
    terminal.UserID = req.UserID
    terminal.FreeRecordBalance = req.FreeRecordBalance

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