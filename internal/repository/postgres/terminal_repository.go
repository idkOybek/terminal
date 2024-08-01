package postgres

import (
	"context"
	"database/sql"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type TerminalRepo struct {
	db *sql.DB
}

func NewTerminalRepository(db *sql.DB) repository.TerminalRepository {
	return &TerminalRepo{db: db}
}

func init() {
    repository.NewTerminalRepository = NewTerminalRepository
}


func (r *TerminalRepo) Create(ctx context.Context, terminal *models.Terminal) error {
	query := `
        INSERT INTO terminals (assembly_number, inn, company_name, address, cash_register_number, status, user_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id`

	err := r.db.QueryRowContext(ctx, query,
		terminal.AssemblyNumber, terminal.INN, terminal.CompanyName, terminal.Address,
		terminal.CashRegisterNumber, terminal.Status, terminal.UserID, terminal.CreatedAt, terminal.UpdatedAt,
	).Scan(&terminal.ID)

	return err
}

func (r *TerminalRepo) GetByID(ctx context.Context, id int) (*models.Terminal, error) {
	query := `
        SELECT id, assembly_number, inn, company_name, address, cash_register_number, status, user_id, created_at, updated_at
        FROM terminals
        WHERE id = $1`

	var terminal models.Terminal
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&terminal.ID, &terminal.AssemblyNumber, &terminal.INN, &terminal.CompanyName, &terminal.Address,
		&terminal.CashRegisterNumber, &terminal.Status, &terminal.UserID, &terminal.CreatedAt, &terminal.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &terminal, nil
}

func (r *TerminalRepo) Update(ctx context.Context, terminal *models.Terminal) error {
	query := `
        UPDATE terminals
        SET inn = $1, company_name = $2, address = $3, cash_register_number = $4, status = $5, updated_at = $6
        WHERE id = $7`

	_, err := r.db.ExecContext(ctx, query,
		terminal.INN, terminal.CompanyName, terminal.Address, terminal.CashRegisterNumber,
		terminal.Status, terminal.UpdatedAt, terminal.ID,
	)

	return err
}

func (r *TerminalRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM terminals WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *TerminalRepo) List(ctx context.Context) ([]*models.Terminal, error) {
	query := `
        SELECT id, assembly_number, inn, company_name, address, cash_register_number, status, user_id, created_at, updated_at
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
			&terminal.ID, &terminal.AssemblyNumber, &terminal.INN, &terminal.CompanyName, &terminal.Address,
			&terminal.CashRegisterNumber, &terminal.Status, &terminal.UserID, &terminal.CreatedAt, &terminal.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		terminals = append(terminals, &terminal)
	}

	return terminals, nil
}

func (r *TerminalRepo) GetByAssemblyNumber(ctx context.Context, assemblyNumber string) (*models.Terminal, error) {
	query := `SELECT id, assembly_number, inn, company_name, address, cash_register_number, status, user_id, created_at, updated_at 
              FROM terminals WHERE assembly_number = $1`

	var terminal models.Terminal
	err := r.db.QueryRowContext(ctx, query, assemblyNumber).Scan(
		&terminal.ID, &terminal.AssemblyNumber, &terminal.INN, &terminal.CompanyName,
		&terminal.Address, &terminal.CashRegisterNumber, &terminal.Status,
		&terminal.UserID, &terminal.CreatedAt, &terminal.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &terminal, nil
}
