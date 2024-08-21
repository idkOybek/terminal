package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
)

type UserRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewUserRepository(db *sql.DB, logger *logger.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
        INSERT INTO users (inn, username, password, company_name, is_active, is_admin)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at`

	err := r.db.QueryRowContext(ctx, query,
		user.INN, user.Username, user.Password, user.CompanyName, user.IsActive, user.IsAdmin,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := `
        SELECT id, inn, username, password, company_name, is_active, is_admin, created_at, updated_at
        FROM users
        WHERE id = $1`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.INN, &user.Username, &user.Password, &user.CompanyName,
		&user.IsActive, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET "
	args := []interface{}{}
	argId := 1

	if user.INN != "" {
		query += fmt.Sprintf("inn = $%d, ", argId)
		args = append(args, user.INN)
		argId++
	}
	if user.Username != "" {
		query += fmt.Sprintf("username = $%d, ", argId)
		args = append(args, user.Username)
		argId++
	}
	if user.Password != "" {
		query += fmt.Sprintf("password = $%d, ", argId)
		args = append(args, user.Password)
		argId++
	}
	if user.CompanyName != "" {
		query += fmt.Sprintf("company_name = $%d, ", argId)
		args = append(args, user.CompanyName)
		argId++
	}
	query += fmt.Sprintf("is_active = $%d, is_admin = $%d, updated_at = $%d ", argId, argId+1, argId+2)
	args = append(args, user.IsActive, user.IsAdmin, time.Now())
	argId += 3

	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf("WHERE id = $%d", argId)
	args = append(args, user.ID)

	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *UserRepository) List(ctx context.Context) ([]*models.User, error) {
	query := `
        SELECT id, inn, username, password, company_name, is_active, is_admin, created_at, updated_at
        FROM users
        ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID, &user.INN, &user.Username, &user.Password, &user.CompanyName,
			&user.IsActive, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
        SELECT id, inn, username, password, is_active, is_admin, created_at, updated_at
        FROM users
        WHERE username = $1`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID, &user.INN, &user.Username, &user.Password,
		&user.IsActive, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
