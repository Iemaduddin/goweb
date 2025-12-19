package service

import (
	"context"

	"github.com/Iemaduddin/goweb/backend-go/internal/model"
	"github.com/Iemaduddin/goweb/backend-go/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	FindAllUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int64) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.userRepo.GetUserByEmail(ctx, email)
}

func (s *userService) FindAllUsers(ctx context.Context) ([]model.User, error) {
	return s.userRepo.FindAllUsers(ctx)
}

func (s *userService) UpdateUser(ctx context.Context, user *model.User) error {
	return s.userRepo.UpdateUser(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return s.userRepo.DeleteUser(ctx, id)
}
