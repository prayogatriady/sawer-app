package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/prayogatriady/sawer-app/model"
	"github.com/prayogatriady/sawer-app/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserService{UserRepository: userRepository}

var (
	user = model.UserEntity{
		ID:        1,
		Username:  "user1",
		Password:  "user1password",
		Balance:   1000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)

func TestUserRepository_GetUser(t *testing.T) {
	ctx := context.Background()

	t.Run("SUCCESS", func(t *testing.T) {
		userRepository.Mock.On("GetUser", ctx, 1).Return(&user, nil)

		result, err := userService.UserRepository.GetUser(ctx, 1)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, user.ID, result.ID)
		assert.Equal(t, user.Username, result.Username)
		assert.Equal(t, user.Password, result.Password)
	})

	t.Run("ERROR", func(t *testing.T) {
		userRepository.Mock.On("GetUser", ctx, 0).Return(nil, errors.New("Error occurred"))

		result, err := userService.UserRepository.GetUser(ctx, 0)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("Error occurred"), err)
	})
}

func TestUserRepository_CreateUser(t *testing.T) {
	ctx := context.Background()

	t.Run("SUCCESS", func(t *testing.T) {
		userRepository.Mock.On("CreateUser", ctx, &user).Return(&user, nil)

		result, err := userService.UserRepository.CreateUser(ctx, &user)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, user.ID, result.ID)
		assert.Equal(t, user.Username, result.Username)
		assert.Equal(t, user.Password, result.Password)
	})

	t.Run("ERROR", func(t *testing.T) {
		userRepository.Mock.On("CreateUser", ctx, &model.UserEntity{}).Return(nil, errors.New("Error occurred"))

		result, err := userService.UserRepository.CreateUser(ctx, &model.UserEntity{})
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("Error occurred"), err)
	})
}

func TestUserRepository_GetUserByUsername(t *testing.T) {
	ctx := context.Background()

	t.Run("SUCCESS", func(t *testing.T) {
		userRepository.Mock.On("GetUserByUsername", ctx, "user1").Return(&user, nil)

		result, err := userService.UserRepository.GetUserByUsername(ctx, "user1")
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, user.ID, result.ID)
		assert.Equal(t, user.Username, result.Username)
		assert.Equal(t, user.Password, result.Password)
	})

	t.Run("ERROR", func(t *testing.T) {
		userRepository.Mock.On("GetUserByUsername", ctx, "").Return(nil, errors.New("Error occurred"))

		result, err := userService.UserRepository.GetUserByUsername(ctx, "")
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("Error occurred"), err)
	})
}

func TestUserRepository_GetUserByUsernamePassword(t *testing.T) {
	ctx := context.Background()

	t.Run("SUCCESS", func(t *testing.T) {
		userRepository.Mock.On("GetUserByUsernamePassword", ctx, "user1", "user1password").Return(&user, nil)

		result, err := userService.UserRepository.GetUserByUsernamePassword(ctx, "user1", "user1password")
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, user.ID, result.ID)
		assert.Equal(t, user.Username, result.Username)
		assert.Equal(t, user.Password, result.Password)
	})

	t.Run("ERROR", func(t *testing.T) {
		userRepository.Mock.On("GetUserByUsernamePassword", ctx, "", "").Return(nil, errors.New("Error occurred"))

		result, err := userService.UserRepository.GetUserByUsernamePassword(ctx, "", "")
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("Error occurred"), err)
	})
}

func TestUserRepository_UpdateUser(t *testing.T) {
	ctx := context.Background()

	t.Run("SUCCESS", func(t *testing.T) {
		userRepository.Mock.On("UpdateUser", ctx, 1, &user).Return(&user, nil)

		result, err := userService.UserRepository.UpdateUser(ctx, 1, &user)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, user.ID, result.ID)
		assert.Equal(t, user.Username, result.Username)
		assert.Equal(t, user.Password, result.Password)
	})

	t.Run("ERROR", func(t *testing.T) {
		userRepository.Mock.On("UpdateUser", ctx, 0, &model.UserEntity{}).Return(nil, errors.New("Error occurred"))

		result, err := userService.UserRepository.UpdateUser(ctx, 0, &model.UserEntity{})
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("Error occurred"), err)
	})
}

func TestUserRepository_DeleteUser(t *testing.T) {
	ctx := context.Background()

	t.Run("SUCCESS", func(t *testing.T) {
		userRepository.Mock.On("DeleteUser", ctx, 1).Return(nil)

		err := userService.UserRepository.DeleteUser(ctx, 1)
		assert.Nil(t, err)
	})

	t.Run("ERROR", func(t *testing.T) {
		userRepository.Mock.On("DeleteUser", ctx, 0).Return(errors.New("Error occurred"))

		err := userService.UserRepository.DeleteUser(ctx, 0)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("Error occurred"), err)
	})
}
