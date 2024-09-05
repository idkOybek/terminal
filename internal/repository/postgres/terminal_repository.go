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
	existingTerminalNumber, existingFiscalModuleNumber, err := r.GetExistingBinding(ctx, terminal.CashRegisterNumber)
	if err != nil {
		return err
	}
	if existingTerminalNumber != "" || existingFiscalModuleNumber != "" {
		if existingTerminalNumber != terminal.CashRegisterNumber || existingFiscalModuleNumber != terminal.ModuleNumber {
			return errors.New("invalid terminal-fiscal module binding")
		}
	}

	query := `
        INSERT INTO terminals (assembly_number, inn, company_name, address, cash_register_number, 
                               module_number, last_request_date, database_update_date, is_active, 
                               user_id, free_record_balance)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        RETURNING id, created_at, updated_at`

	err = r.db.QueryRowContext(ctx, query,
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
               user_id, free_record_balance, created_at, updated_at, status_changed_by_admin
        FROM terminals
        WHERE id = $1`

	var terminal models.Terminal
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&terminal.ID, &terminal.AssemblyNumber, &terminal.INN, &terminal.CompanyName,
		&terminal.Address, &terminal.CashRegisterNumber, &terminal.ModuleNumber,
		&terminal.LastRequestDate, &terminal.DatabaseUpdateDate, &terminal.IsActive,
		&terminal.UserID, &terminal.FreeRecordBalance, &terminal.CreatedAt, &terminal.UpdatedAt,
		&terminal.StatusChangedByAdmin,
	)

	if err != nil {
		return nil, err
	}

	return &terminal, nil
}

func (r *TerminalRepository) Update(ctx context.Context, terminal *models.Terminal) error {
	existingTerminalNumber, existingFiscalModuleNumber, err := r.GetExistingBinding(ctx, terminal.CashRegisterNumber)
	if err != nil {
		return err
	}
	if existingTerminalNumber != "" || existingFiscalModuleNumber != "" {
		if (existingTerminalNumber != terminal.CashRegisterNumber && existingTerminalNumber != "") ||
			(existingFiscalModuleNumber != terminal.ModuleNumber && existingFiscalModuleNumber != "") {
			return errors.New("invalid terminal-fiscal module binding")
		}
	}
	query := "UPDATE terminals SET "
	args := []interface{}{}
	argId := 1

	// Функция для добавления поля в запрос, если оно не пустое
	addField := func(field string, value interface{}) {
		if value != nil && value != "" {
			query += fmt.Sprintf("%s = $%d, ", field, argId)
			args = append(args, value)
			argId++
		}
	}

	// Добавляем поля в запрос, только если они не пустые
	addField("assembly_number", terminal.AssemblyNumber)
	addField("inn", terminal.INN)
	addField("company_name", terminal.CompanyName)
	addField("address", terminal.Address)
	addField("cash_register_number", terminal.CashRegisterNumber)
	addField("module_number", terminal.ModuleNumber)
	addField("last_request_date", terminal.LastRequestDate)
	addField("database_update_date", terminal.DatabaseUpdateDate)
	addField("is_active", terminal.IsActive)
	addField("free_record_balance", terminal.FreeRecordBalance)
	addField("status_changed_by_admin", terminal.StatusChangedByAdmin)

	// Всегда обновляем поле updated_at
	query += fmt.Sprintf("updated_at = $%d ", argId)
	args = append(args, time.Now())
	argId++

	// Удаляем последнюю запятую, если она есть
	query = strings.TrimSuffix(query, ", ")

	// Добавляем условие WHERE
	query += fmt.Sprintf("WHERE id = $%d", argId)
	args = append(args, terminal.ID)

	// Логируем запрос и аргументы
	r.logger.Info("Updating terminal",
		"query", query,
		"args", fmt.Sprintf("%+v", args))

	// Выполняем запрос
	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update terminal: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	r.logger.Info("Terminal updated successfully",
		"id", terminal.ID,
		"rows_affected", rowsAffected)

	return nil
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
               user_id, free_record_balance, created_at, updated_at, status_changed_by_admin
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
			&terminal.StatusChangedByAdmin,
		)
		if err != nil {
			return nil, err
		}
		terminals = append(terminals, &terminal)
	}

	return terminals, nil
}

func (r *TerminalRepository) CheckTerminalFiscalModuleBinding(ctx context.Context, terminalNumber, fiscalModuleNumber string) (bool, error) {
	query := `
        SELECT COUNT(*) 
        FROM terminals 
        WHERE cash_register_number = $1 AND module_number = $2
    `
	var count int
	err := r.db.QueryRowContext(ctx, query, terminalNumber, fiscalModuleNumber).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking terminal-fiscal module binding: %w", err)
	}
	return count > 0, nil
}

func (r *TerminalRepository) GetExistingBinding(ctx context.Context, number string) (string, string, error) {
	query := `
        SELECT cash_register_number, module_number 
        FROM terminals 
        WHERE cash_register_number = $1 OR module_number = $1
    `
	var terminalNumber, fiscalModuleNumber string
	err := r.db.QueryRowContext(ctx, query, number).Scan(&terminalNumber, &fiscalModuleNumber)
	if err == sql.ErrNoRows {
		return "", "", nil
	}
	if err != nil {
		return "", "", fmt.Errorf("error getting existing binding: %w", err)
	}
	return terminalNumber, fiscalModuleNumber, nil
}
