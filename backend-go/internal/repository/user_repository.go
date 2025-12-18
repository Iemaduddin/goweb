package repository

import (
	"context"
	"database/sql"

	"github.com/Iemaduddin/goweb/backend-go/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Content, user *model.User) error
	GetUserByID(ctx context.Content, id int64) (*model.User, error)
	GetUserByEmail(ctx context.Content, email string) (*model.User, error)
	FindAllUsers(ctx context.Content) ([]model.User, error)
	UpdateUser(ctx context.Content, user *model.User) error
	DeleteUser(ctx context.Content, id int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Content, user *model.User) error {
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
		return "Terjadi kesalahan: " + err.Error()
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "Terjadi kesalahan: " + err.Error()
	}

	user.ID = id

	return nil
}
