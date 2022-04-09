package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/NwokoyeChigozie/quik_task/internal/model"
	"github.com/NwokoyeChigozie/quik_task/internal/user"
	"github.com/NwokoyeChigozie/quik_task/pkg/mocks/repository"
	"github.com/stretchr/testify/assert"
)

func Test_CreateUser(t *testing.T) {
	ast := assert.New(t)
	t.Run("testing Create User", func(t *testing.T) {
		testData := model.User{
			ID:           1,
			FirstName:    "firstname",
			LastName:     "lastname",
			Email:        "email@email.com",
			Password:     "password",
			Token:        "token",
			TokenExpires: time.Now(),
			WalletID:     1,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		requestData := model.CreateUserRequestModel{
			FirstName: "firstname",
			LastName:  "lastname",
			Email:     "email@email.com",
			Password:  "password",
		}
		userRepository := &repository.User{TestData: &testData}

		userRepository.On("CreateUser", context.Background(), requestData).Return(testData, 0, nil).Once()

		service := user.NewUserService(userRepository)
		user, code, err := service.CreateUser(context.Background(), requestData)
		ast.NoError(err)
		ast.Equal(user, &testData)
		ast.Equal(code, 0)

		userRepository.AssertExpectations(t)
	})
}
func Test_GetUser(t *testing.T) {
	ast := assert.New(t)
	t.Run("testing Get User", func(t *testing.T) {
		testData := *&model.User{
			ID:           1,
			FirstName:    "firstname",
			LastName:     "lastname",
			Email:        "email@email.com",
			Password:     "password",
			Token:        "token",
			TokenExpires: time.Now(),
			WalletID:     1,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		userRepository := &repository.User{TestData: &testData}

		userRepository.On("GetUser", context.Background(), int(testData.ID)).Return(testData, 0, nil).Once()

		service := user.NewUserService(userRepository)
		user, code, err := service.GetUser(context.Background(), int(testData.ID))
		ast.NoError(err)
		ast.Equal(user, &testData)
		ast.Equal(code, 0)

		userRepository.AssertExpectations(t)
	})
}

func Test_Login(t *testing.T) {
	ast := assert.New(t)
	t.Run("testing Login User", func(t *testing.T) {
		testData := model.User{
			ID:           1,
			FirstName:    "firstname",
			LastName:     "lastname",
			Email:        "email@email.com",
			Password:     "password",
			Token:        "token",
			TokenExpires: time.Now(),
			WalletID:     1,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		req := model.LoginRequestModel{
			Email:    "email@email.com",
			Password: "password"}
		userRepository := &repository.User{TestData: &testData}

		userRepository.On("Login", context.Background(), req).Return(testData, 0, nil).Once()

		service := user.NewUserService(userRepository)
		user, code, err := service.Login(context.Background(), req)
		ast.NoError(err)
		ast.Equal(user, &testData)
		ast.Equal(code, 0)

		userRepository.AssertExpectations(t)
	})
}
