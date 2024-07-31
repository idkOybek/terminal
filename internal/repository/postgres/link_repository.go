// internal/repository/postgres/link_repository.go

package postgres

import (
	"context"
	"database/sql"

	"github.com/idkOybek/newNewTerminal/internal/models"
)

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) Create(ctx context.Context, link *models.Link) error {
	query := `
		INSERT INTO links (fiscal_number, factory_number)
		VALUES ($1, $2)
		RETURNING id, created_at`
	
	err := r.db.QueryRowContext(ctx, query,
		link.FiscalNumber, link.FactoryNumber,
	).Scan(&link.ID, &link.CreatedAt)

	return err
}

func (r *LinkRepository) GetByFactoryNumber(ctx context.Context, factoryNumber string) (*models.Link, error) {
	query := `
		SELECT id, fiscal_number, factory_number, created_at
		FROM links
		WHERE factory_number = $1`

	var link models.Link
	err := r.db.QueryRowContext(ctx, query, factoryNumber).Scan(
		&link.ID, &link.FiscalNumber, &link.FactoryNumber, &link.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *LinkRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM links WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *LinkRepository) List(ctx context.Context) ([]models.Link, error) {
	query := `
		SELECT id, fiscal_number, factory_number, created_at
		FROM links
		ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []models.Link
	for rows.Next() {
		var link models.Link
		err := rows.Scan(
			&link.ID, &link.FiscalNumber, &link.FactoryNumber, &link.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	return links, nil
}