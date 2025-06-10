package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AgsOtero/event-ticket-api/internal/core/domain"
	"github.com/AgsOtero/event-ticket-api/internal/core/ports"
)

type postgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) ports.UserRepository {
	return &postgresUserRepository{
		db: db,
	}
}

func (r *postgresUserRepository) Save(ctx context.Context, user domain.User) (domain.User, error) {
	query := `INSERT INTO users (name, surname, email, phone, password_hash) 
              VALUES ($1, $2, $3, $4, $5) 
              RETURNING id`

	err := r.db.QueryRowContext(ctx, query, user.Name, user.Surname, user.Email, user.Phone, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *postgresUserRepository) FindByID(ctx context.Context, id int64) (domain.User, error) {
	query := `SELECT id, name, surname, email, phone, password_hash, created_at 
              FROM users 
              WHERE id = $1`

	var user domain.User
	row := r.db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *postgresUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	query := `SELECT id, name, surname, email, phone, password_hash FROM users WHERE email = $1`
	var user domain.User
	row := r.db.QueryRowContext(ctx, query, email)

	err :=
		row.Scan(&user.ID,
			&user.Name,
			&user.Surname,
			&user.Email,
			&user.PasswordHash,
			&user.Phone,
		)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}
	return user, nil
}
