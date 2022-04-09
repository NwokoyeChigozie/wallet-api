package repository

import (
	"context"

	"github.com/NwokoyeChigozie/quik_task/internal/model"
	"github.com/stretchr/testify/mock"
)

type User struct {
	mock.Mock
	TestData *model.User
}

func (u *User) CreateUser(ctx context.Context, form model.CreateUserRequestModel) (*model.User, int, error) {
	args := u.Called(ctx, form)
	return u.TestData, args.Int(1), args.Error(2)
}

func (u *User) GetUser(ctx context.Context, userID int) (*model.User, int, error) {
	args := u.Called(ctx, userID)
	return u.TestData, args.Int(1), args.Error(2)
}

func (u *User) Login(ctx context.Context, form model.LoginRequestModel) (*model.User, int, error) {
	args := u.Called(ctx, form)
	return u.TestData, args.Int(1), args.Error(2)
}
