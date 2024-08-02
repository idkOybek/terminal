package postgres

import (
    "context"
    "database/sql"
    "github.com/idkOybek/newNewTerminal/internal/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
    query := `
        INSERT INTO users (inn, username, password, is_active, is_admin)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, updated_at`
    
    err := r.db.QueryRowContext(ctx, query,
        user.INN, user.Username, user.Password, user.IsActive, user.IsAdmin,
    ).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

    return err
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
    query := `
        SELECT id, inn, username, password, is_active, is_admin, created_at, updated_at
        FROM users
        WHERE id = $1`

    var user models.User
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &user.ID, &user.INN, &user.Username, &user.Password,
        &user.IsActive, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
    )

    if err != nil {
        return nil, err
    }

    return &user, nil
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

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
    query := `
        UPDATE users
        SET inn = $1, username = $2, password = $3, is_active = $4, is_admin = $5, updated_at = NOW()
        WHERE id = $6`

    _, err := r.db.ExecContext(ctx, query,
        user.INN, user.Username, user.Password, user.IsActive, user.IsAdmin, user.ID,
    )

    return err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
    query := `DELETE FROM users WHERE id = $1`

    _, err := r.db.ExecContext(ctx, query, id)

    return err
}

func (r *UserRepository) List(ctx context.Context) ([]*models.User, error) {
    query := `
        SELECT id, inn, username, password, is_active, is_admin, created_at, updated_at
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
            &user.ID, &user.INN, &user.Username, &user.Password,
            &user.IsActive, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        users = append(users, &user)
    }

    return users, nil
}