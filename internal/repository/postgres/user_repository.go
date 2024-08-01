package postgres

import (
	"context"
	"database/sql"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/repository"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepo{db: db}
}

func init() {
	repository.NewUserRepository = NewUserRepository
}

func (r *UserRepo) Create(ctx context.Context, user *models.User) error {
	query := `
        INSERT INTO users (username, password, is_admin, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`

	err := r.db.QueryRowContext(ctx, query,
		user.Username, user.Password, user.IsAdmin, user.CreatedAt, user.UpdatedAt,
	).Scan(&user.ID)

	return err
}

func (r *UserRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := `
        SELECT id, username, password, is_admin, created_at, updated_at
        FROM users
        WHERE id = $1`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Username, &user.Password, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
        SELECT id, username, password, is_admin, created_at, updated_at
        FROM users
        WHERE username = $1`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID, &user.Username, &user.Password, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *models.User) error {
	query := `
        UPDATE users
        SET username = $1, password = $2, is_admin = $3, updated_at = $4
        WHERE id = $5`

	_, err := r.db.ExecContext(ctx, query,
		user.Username, user.Password, user.IsAdmin, user.UpdatedAt, user.ID,
	)

	return err
}

func (r *UserRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *UserRepo) List(ctx context.Context) ([]*models.User, error) {
	query := `
        SELECT id, username, is_admin, created_at, updated_at
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
			&user.ID, &user.Username, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
