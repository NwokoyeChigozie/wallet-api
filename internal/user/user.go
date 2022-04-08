package user

import (
	"context"

	"github.com/NwokoyeChigozie/quik_task/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, form model.CreateUserRequestModel) (*model.User, int, error)
	GetUser(ctx context.Context, userID int) (*model.User, int, error)
	Login(ctx context.Context, form model.LoginRequestModel) (*model.User, int, error)
}

type Service interface {
	CreateUser(ctx context.Context, form model.CreateUserRequestModel) (*model.User, int, error)
	GetUser(ctx context.Context, userID int) (*model.User, int, error)
	Login(ctx context.Context, form model.LoginRequestModel) (*model.User, int, error)
}

type user struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) Service {
	return &user{userRepository: userRepository}
}

func (p *user) CreateUser(ctx context.Context, form model.CreateUserRequestModel) (*model.User, int, error) {
	return p.userRepository.CreateUser(ctx, form)
}

func (p *user) GetUser(ctx context.Context, userID int) (*model.User, int, error) {
	return p.userRepository.GetUser(ctx, userID)
}

func (p *user) Login(ctx context.Context, form model.LoginRequestModel) (*model.User, int, error) {
	return p.userRepository.Login(ctx, form)
}
