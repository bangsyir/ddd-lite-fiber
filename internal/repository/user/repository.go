package repository

import (
	"DDD-fiberv2/internal/domain/user"
	"context"
	"database/sql"
	"errors"
)

type sqliteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) user.Repository {
	return &sqliteRepository{db: db}
}

func (r *sqliteRepository) Create(ctx context.Context, user *user.User) error {
	query := `
  INSERT INTO users (id, name, email, created_at, updated_at)
  VALUES (?,?,?,?,?)
  `
	_, err := r.db.ExecContext(
		ctx,
		query,
		user.ID,
		user.Name,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		if err.Error() == "UNIQUE contraint failed: users.email" {
			return errors.New("email already exist")
		}
		return err
	}
	return nil
}

func (r *sqliteRepository) FindById(ctx context.Context, id string) (*user.User, error) {
	query := `
  SELECT id, name, email, created_at, updated_at
  FROM users
  WHERE id = ?
  `
	user := &user.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *sqliteRepository) Update(ctx context.Context, user *user.User) error {
	query := `
  UPDATE users
  SET name = ?, email = ?, updated_at = ?
  WHERE id = ?
  `
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
			return errors.New("email already exists")
		}
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
func (r *sqliteRepository) Delete(ctx context.Context, id string) error {

	query := `DELETE FROM users WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *sqliteRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	query := `
  SELECT id, name, email, created_at, updated_at
  FROM users
  `
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*user.User{}

	for rows.Next() {
		user := &user.User{}
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil

}
