package service

import (
	"context"
	"testing"
	"time"

	"github.com/prayogatriady/sawer-app/model"
	"github.com/prayogatriady/sawer-app/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserService{UserRepository: userRepository}

var user = model.UserEntity{
	ID:        1,
	Username:  "user1",
	Password:  "user1password",
	Balance:   1000,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestUserRepository_GetUser(t *testing.T) {
	ctx := context.Background()

	userRepository.Mock.On("GetUser", ctx, 1).Return(user)

	result, err := userService.UserRepository.GetUser(ctx, 1)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Username, result.Username)
	assert.Equal(t, user.Password, result.Password)
}
