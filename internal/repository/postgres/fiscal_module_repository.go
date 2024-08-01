package postgres

import (
	"context"
	"database/sql"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type FiscalModuleRepo struct {
	db *sql.DB
}

func NewFiscalModuleRepository(db *sql.DB) repository.FiscalModuleRepository {
	return &FiscalModuleRepo{db: db}
}

func init() {
	repository.NewFiscalModuleRepository = NewFiscalModuleRepository
}

func (r *FiscalModuleRepo) Create(ctx context.Context, module *models.FiscalModule) error {
	query := `
        INSERT INTO fiscal_modules (fiscal_number, factory_number, user_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`

	err := r.db.QueryRowContext(ctx, query,
		module.FiscalNumber, module.FactoryNumber, module.UserID, module.CreatedAt, module.UpdatedAt,
	).Scan(&module.ID)

	return err
}

func (r *FiscalModuleRepo) GetByID(ctx context.Context, id int) (*models.FiscalModule, error) {
	query := `
        SELECT id, fiscal_number, factory_number, user_id, created_at, updated_at
        FROM fiscal_modules
        WHERE id = $1`

	var module models.FiscalModule
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&module.ID, &module.FiscalNumber, &module.FactoryNumber,
		&module.UserID, &module.CreatedAt, &module.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &module, nil
}

func (r *FiscalModuleRepo) GetByFactoryNumber(ctx context.Context, factoryNumber string) (*models.FiscalModule, error) {
	query := `
        SELECT id, fiscal_number, factory_number, user_id, created_at, updated_at
        FROM fiscal_modules
        WHERE factory_number = $1`

	var module models.FiscalModule
	err := r.db.QueryRowContext(ctx, query, factoryNumber).Scan(
		&module.ID, &module.FiscalNumber, &module.FactoryNumber,
		&module.UserID, &module.CreatedAt, &module.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &module, nil
}

func (r *FiscalModuleRepo) Update(ctx context.Context, module *models.FiscalModule) error {
	query := `
        UPDATE fiscal_modules
        SET fiscal_number = $1, factory_number = $2, user_id = $3, updated_at = $4
        WHERE id = $5`

	_, err := r.db.ExecContext(ctx, query,
		module.FiscalNumber, module.FactoryNumber, module.UserID, module.UpdatedAt, module.ID,
	)

	return err
}

func (r *FiscalModuleRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM fiscal_modules WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *FiscalModuleRepo) List(ctx context.Context) ([]*models.FiscalModule, error) {
	query := `
        SELECT id, fiscal_number, factory_number, user_id, created_at, updated_at
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
			&module.UserID, &module.CreatedAt, &module.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		modules = append(modules, &module)
	}

	return modules, nil
}
