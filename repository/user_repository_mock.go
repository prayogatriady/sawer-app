package repository

import (
	"context"
	"errors"

	"github.com/prayogatriady/sawer-app/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (repository *UserRepositoryMock) GetUser(ctx context.Context, userID int) (*model.UserEntity, error) {
	arguments := repository.Mock.Called(ctx, userID)
	if arguments.Get(0) == nil {
		return nil, errors.New("error get user")
	} else {
		user := arguments.Get(0).(model.UserEntity)
		return &user, nil
	}
}

func (repository *UserRepositoryMock) CreateUser(ctx context.Context, user *model.UserEntity) (*model.UserEntity, error) {
	// arguments := repository.Mock.Called(ctx, user)
	return &model.UserEntity{}, nil
}

func (repository *UserRepositoryMock) GetUserByUsername(ctx context.Context, username string) (*model.UserEntity, error) {
	return &model.UserEntity{}, nil
}

func (repository *UserRepositoryMock) GetUserByUsernamePassword(ctx context.Context, username string, password string) (*model.UserEntity, error) {
	return &model.UserEntity{}, nil
}

func (repository *UserRepositoryMock) UpdateUser(ctx context.Context, userID int, updateUser *model.UserEntity) (*model.UserEntity, error) {
	return &model.UserEntity{}, nil
}

func (repository *UserRepositoryMock) DeleteUser(ctx context.Context, userID int) error {
	return nil
}
