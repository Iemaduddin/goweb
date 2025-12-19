package repository

import (
	"context"
	"database/sql"

	"github.com/Iemaduddin/goweb/backend-go/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	FindAllUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (name, email, password, role, created_at)
		VALUES (?, ?, ?, ?, ?)
	`

	res, err := r.db.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.CreatedAt,
	)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id

	return nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	query := `
		SELECT id, name, email, password, role, created_at
		FROM users
		WHERE id = ?
	`

	row := r.db.QueryRowContext(ctx, query, id)

	user := &model.User{}

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT id, name, email, password, role, created_at
		FROM users
		WHERE email = ?
	`

	row := r.db.QueryRowContext(ctx, query, email)

	user := &model.User{}

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindAllUsers(ctx context.Context) ([]model.User, error) {
	query := `
		SELECT id, name, email, password, role, created_at
		FROM users
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User

	for rows.Next() {
		user := &model.User{}

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return users, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	query := `
		UPDATE users
		SET name = ?, email = ?, password = ?, role = ?, created_at = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.CreatedAt,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	query := `
		DELETE FROM users
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
