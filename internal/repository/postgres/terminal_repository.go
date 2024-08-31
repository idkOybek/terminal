package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
)

type TerminalRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewTerminalRepository(db *sql.DB, logger *logger.Logger) *TerminalRepository {
	return &TerminalRepository{
		db:     db,
		logger: logger,
	}
}

func (r *TerminalRepository) GetByCashRegisterNumber(ctx context.Context, cashRegisterNumber string) (*models.Terminal, error) {
	var terminal models.Terminal
	err := r.db.QueryRowContext(ctx, `
        SELECT id, assembly_number, inn, company_name, address, cash_register_number, 
               module_number, last_request_date, database_update_date, is_active, user_id, 
               free_record_balance, created_at, updated_at 
        FROM terminals 
        WHERE cash_register_number = $1
    `, cashRegisterNumber).Scan(
		&terminal.ID, &terminal.AssemblyNumber, &terminal.INN, &terminal.CompanyName,
		&terminal.Address, &terminal.CashRegisterNumber, &terminal.ModuleNumber,
		&terminal.LastRequestDate, &terminal.DatabaseUpdateDate, &terminal.IsActive,
		&terminal.UserID, &terminal.FreeRecordBalance, &terminal.CreatedAt, &terminal.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &terminal, nil
}

func (r *TerminalRepository) GetStatus(ctx context.Context, id int) (bool, error) {
	var isActive bool
	err := r.db.QueryRowContext(ctx, `
        SELECT is_active FROM terminals WHERE id = $1
    `, id).Scan(&isActive)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.New("terminal not found")
		}
		return false, err
	}
	return isActive, nil
}

func (r *TerminalRepository) GetUserIDByCashRegisterNumber(ctx context.Context, cashRegisterNumber string) (int, error) {
	query := `
		SELECT user_id 
		FROM fiscal_modules
		WHERE factory_number = $1`

	var userID int
	err := r.db.QueryRowContext(ctx, query, cashRegisterNumber).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no user associated with factory number %s", cashRegisterNumber)
		}
		return 0, err
	}

	return userID, nil
}

func (r *TerminalRepository) Create(ctx context.Context, terminal *models.Terminal) error {
	query := `
        INSERT INTO terminals (assembly_number, inn, company_name, address, cash_register_number, 
                               module_number, last_request_date, database_update_date, is_active, 
                               user_id, free_record_balance)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        RETURNING id, created_at, updated_at`

	err := r.db.QueryRowContext(ctx, query,
		terminal.AssemblyNumber, terminal.INN, terminal.CompanyName, terminal.Address,
		terminal.CashRegisterNumber, terminal.ModuleNumber, terminal.LastRequestDate,
		terminal.DatabaseUpdateDate, terminal.IsActive, terminal.UserID, terminal.FreeRecordBalance,
	).Scan(&terminal.ID, &terminal.CreatedAt, &terminal.UpdatedAt)

	return err
}

func (r *TerminalRepository) GetByID(ctx context.Context, id int) (*models.Terminal, error) {
	query := `
        SELECT id, assembly_number, inn, company_name, address, cash_register_number, 
               module_number, last_request_date, database_update_date, is_active, 
               user_id, free_record_balance, created_at, updated_at
        FROM terminals
        WHERE id = $1`

	var terminal models.Terminal
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&terminal.ID, &terminal.AssemblyNumber, &terminal.INN, &terminal.CompanyName,
		&terminal.Address, &terminal.CashRegisterNumber, &terminal.ModuleNumber,
		&terminal.LastRequestDate, &terminal.DatabaseUpdateDate, &terminal.IsActive,
		&terminal.UserID, &terminal.FreeRecordBalance, &terminal.CreatedAt, &terminal.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &terminal, nil
}

func (r *TerminalRepository) Update(ctx context.Context, terminal *models.Terminal) error {
	query := "UPDATE terminals SET "
	args := []interface{}{}
	argId := 1

	// Динамически формируем запрос для всех полей
	if terminal.AssemblyNumber != "" {
		query += fmt.Sprintf("assembly_number = $%d, ", argId)
		args = append(args, terminal.AssemblyNumber)
		argId++
	}
	if terminal.INN != "" {
		query += fmt.Sprintf("inn = $%d, ", argId)
		args = append(args, terminal.INN)
		argId++
	}
	if terminal.CompanyName != "" {
		query += fmt.Sprintf("company_name = $%d, ", argId)
		args = append(args, terminal.CompanyName)
		argId++
	}
	if terminal.Address != "" {
		query += fmt.Sprintf("address = $%d, ", argId)
		args = append(args, terminal.Address)
		argId++
	}
	if terminal.CashRegisterNumber != "" {
		query += fmt.Sprintf("cash_register_number = $%d, ", argId)
		args = append(args, terminal.CashRegisterNumber)
		argId++
	}
	if terminal.ModuleNumber != "" {
		query += fmt.Sprintf("module_number = $%d, ", argId)
		args = append(args, terminal.ModuleNumber)
		argId++
	}

	// Добавляем обновление для LastRequestDate и DatabaseUpdateDate
	query += fmt.Sprintf("last_request_date = $%d, ", argId)
	args = append(args, terminal.LastRequestDate)
	argId++

	query += fmt.Sprintf("database_update_date = $%d, ", argId)
	args = append(args, terminal.DatabaseUpdateDate)
	argId++

	// Всегда обновляем is_active
	query += fmt.Sprintf("is_active = $%d, ", argId)
	args = append(args, terminal.IsActive)
	argId++

	query += fmt.Sprintf("user_id = $%d, ", argId)
	args = append(args, terminal.UserID)
	argId++

	query += fmt.Sprintf("free_record_balance = $%d, ", argId)
	args = append(args, terminal.FreeRecordBalance)
	argId++

	// Добавляем новое поле StatusChangedByAdmin
	query += fmt.Sprintf("status_changed_by_admin = $%d, ", argId)
	args = append(args, terminal.StatusChangedByAdmin)
	argId++

	// Всегда обновляем поле updated_at
	query += fmt.Sprintf("updated_at = $%d ", argId)
	args = append(args, time.Now())
	argId++

	// Удаляем последнюю запятую и пробел, если они есть
	query = strings.TrimSuffix(query, ", ")

	query += fmt.Sprintf("WHERE id = $%d", argId)
	args = append(args, terminal.ID)

	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *TerminalRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM terminals WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *TerminalRepository) List(ctx context.Context) ([]*models.Terminal, error) {
	query := `
        SELECT id, assembly_number, inn, company_name, address, cash_register_number, 
               module_number, last_request_date, database_update_date, is_active, 
               user_id, free_record_balance, created_at, updated_at
        FROM terminals
        ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var terminals []*models.Terminal
	for rows.Next() {
		var terminal models.Terminal
		err := rows.Scan(
			&terminal.ID, &terminal.AssemblyNumber, &terminal.INN, &terminal.CompanyName,
			&terminal.Address, &terminal.CashRegisterNumber, &terminal.ModuleNumber,
			&terminal.LastRequestDate, &terminal.DatabaseUpdateDate, &terminal.IsActive,
			&terminal.UserID, &terminal.FreeRecordBalance, &terminal.CreatedAt, &terminal.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		terminals = append(terminals, &terminal)
	}

	return terminals, nil
}
