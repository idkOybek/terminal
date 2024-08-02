package postgres

import (
    "context"
    "database/sql"
    "github.com/idkOybek/newNewTerminal/internal/models"
)

type FiscalModuleRepository struct {
    db *sql.DB
}

func NewFiscalModuleRepository(db *sql.DB) *FiscalModuleRepository {
    return &FiscalModuleRepository{db: db}
}

func (r *FiscalModuleRepository) Create(ctx context.Context, module *models.FiscalModule) error {
    query := `
        INSERT INTO fiscal_modules (fiscal_number, factory_number, user_id)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at`
    
    err := r.db.QueryRowContext(ctx, query,
        module.FiscalNumber, module.FactoryNumber, module.UserID,
    ).Scan(&module.ID, &module.CreatedAt, &module.UpdatedAt)

    return err
}

func (r *FiscalModuleRepository) GetByID(ctx context.Context, id int) (*models.FiscalModule, error) {
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

func (r *FiscalModuleRepository) Update(ctx context.Context, module *models.FiscalModule) error {
    query := `
        UPDATE fiscal_modules
        SET fiscal_number = $1, factory_number = $2, user_id = $3, updated_at = NOW()
        WHERE id = $4`

    _, err := r.db.ExecContext(ctx, query,
        module.FiscalNumber, module.FactoryNumber, module.UserID, module.ID,
    )

    return err
}

func (r *FiscalModuleRepository) Delete(ctx context.Context, id int) error {
    query := `DELETE FROM fiscal_modules WHERE id = $1`

    _, err := r.db.ExecContext(ctx, query, id)

    return err
}

func (r *FiscalModuleRepository) List(ctx context.Context) ([]*models.FiscalModule, error) {
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