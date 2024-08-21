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

type FiscalModuleRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewFiscalModuleRepository(db *sql.DB, logger *logger.Logger) *FiscalModuleRepository {
	return &FiscalModuleRepository{
		db:     db,
		logger: logger,
	}
}

func (r *FiscalModuleRepository) Create(ctx context.Context, module *models.FiscalModule) error {
	query := `
        INSERT INTO fiscal_modules (fiscal_number, factory_number, user_id, is_active)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at`

	err := r.db.QueryRowContext(ctx, query,
		module.FiscalNumber, module.FactoryNumber, module.UserID, module.IsActive,
	).Scan(&module.ID, &module.CreatedAt, &module.UpdatedAt)

	return err
}

func (r *FiscalModuleRepository) GetByFactoryNumber(ctx context.Context, factoryNumber string) (*models.FiscalModule, error) {
	query := `SELECT id, fiscal_number, factory_number, user_id, created_at, updated_at FROM fiscal_modules WHERE factory_number = $1`

	var module models.FiscalModule
	err := r.db.QueryRowContext(ctx, query, factoryNumber).Scan(
		&module.ID, &module.FiscalNumber, &module.FactoryNumber,
		&module.UserID, &module.CreatedAt, &module.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &module, nil
}

func (r *FiscalModuleRepository) GetByID(ctx context.Context, id int) (*models.FiscalModule, error) {
	query := `
        SELECT id, fiscal_number, factory_number, user_id, is_active, created_at, updated_at
        FROM fiscal_modules
        WHERE id = $1`

	var module models.FiscalModule
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&module.ID, &module.FiscalNumber, &module.FactoryNumber,
		&module.UserID, &module.IsActive, &module.CreatedAt, &module.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &module, nil
}

func (r *FiscalModuleRepository) Update(ctx context.Context, module *models.FiscalModule) error {
	r.logger.Info("Starting fiscal module update", "id", module.ID, "is_active", module.IsActive)

	query := "UPDATE fiscal_modules SET "
	args := []interface{}{}
	argId := 1

	if module.FiscalNumber != "" {
		query += fmt.Sprintf("fiscal_number = $%d, ", argId)
		args = append(args, module.FiscalNumber)
		argId++
	}
	if module.FactoryNumber != "" {
		query += fmt.Sprintf("factory_number = $%d, ", argId)
		args = append(args, module.FactoryNumber)
		argId++
	}
	query += fmt.Sprintf("user_id = $%d, is_active = $%d, updated_at = $%d ", argId, argId+1, argId+2)
	args = append(args, module.UserID, module.IsActive, time.Now())
	argId += 3

	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf("WHERE id = $%d", argId)
	args = append(args, module.ID)

	r.logger.Info("Executing update query", "query", query, "args", args)

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		r.logger.Error("Failed to execute update query", "error", err)
		return fmt.Errorf("failed to update fiscal module: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Error("Failed to get rows affected", "error", err)
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	r.logger.Info("Fiscal module update completed", "id", module.ID, "rows_affected", rowsAffected)

	if rowsAffected == 0 {
		r.logger.Warn("No rows were updated", "id", module.ID)
		return errors.New("no rows were updated")
	}

	return nil
}
func (r *FiscalModuleRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM fiscal_modules WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *FiscalModuleRepository) List(ctx context.Context) ([]*models.FiscalModule, error) {
	query := `
        SELECT id, fiscal_number, factory_number, user_id, is_active, created_at, updated_at
        FROM fiscal_modules
        ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []*models.FiscalModule
	for rows.Next() {
		var module models.FiscalModule
		err := rows.Scan(
			&module.ID, &module.FiscalNumber, &module.FactoryNumber,
			&module.UserID, &module.IsActive, &module.CreatedAt, &module.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		modules = append(modules, &module)
	}

	return modules, nil
}

func (r *FiscalModuleRepository) DeleteByUserID(ctx context.Context, userID int) error {
	query := `DELETE FROM fiscal_modules WHERE user_id = $1`

	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}
