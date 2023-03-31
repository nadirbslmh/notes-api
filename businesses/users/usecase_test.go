package users_test

import (
	"context"
	"errors"
	"notes-api/app/middlewares"
	"notes-api/businesses/users"
	_userMock "notes-api/businesses/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	usersRepository _userMock.Repository
	usersService    users.Usecase

	usersDomain users.Domain
	ctx         context.Context
)

func TestMain(m *testing.M) {
	usersService = users.NewUserUseCase(&usersRepository, &middlewares.JWTConfig{})

	usersDomain = users.Domain{
		Email:    "test@test.com",
		Password: "123123",
	}

	ctx = context.TODO()

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		usersRepository.On("Register", ctx, &usersDomain).Return(usersDomain, nil).Once()

		result, err := usersService.Register(ctx, &usersDomain)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Register | InValid", func(t *testing.T) {
		usersRepository.On("Register", ctx, &users.Domain{}).Return(users.Domain{}, errors.New("failed")).Once()

		result, err := usersService.Register(ctx, &users.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		usersRepository.On("GetByEmail", ctx, &usersDomain).Return(users.Domain{}, nil).Once()

		result, err := usersService.Login(ctx, &usersDomain)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		usersRepository.On("GetByEmail", ctx, &users.Domain{}).Return(users.Domain{}, errors.New("failed")).Once()

		result, err := usersService.Login(ctx, &users.Domain{})

		assert.Empty(t, result)
		assert.NotNil(t, err)
	})
}
