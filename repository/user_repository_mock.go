package repository

import (
	"context"

	"github.com/prayogatriady/sawer-app/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (r *UserRepositoryMock) GetUser(ctx context.Context, userID int) (*model.UserEntity, error) {
	arguments := r.Mock.Called(ctx, userID)

	var (
		return1 *model.UserEntity
		return2 error
	)

	if arguments.Get(0) != nil {
		return1 = arguments.Get(0).(*model.UserEntity)
	}
	if arguments.Get(1) != nil {
		return2 = arguments.Get(1).(error)
	}

	return return1, return2
}

func (r *UserRepositoryMock) CreateUser(ctx context.Context, user *model.UserEntity) (*model.UserEntity, error) {
	arguments := r.Mock.Called(ctx, user)

	var (
		return1 *model.UserEntity
		return2 error
	)

	if arguments.Get(0) != nil {
		return1 = arguments.Get(0).(*model.UserEntity)
	}
	if arguments.Get(1) != nil {
		return2 = arguments.Get(1).(error)
	}

	return return1, return2
}

func (r *UserRepositoryMock) GetUserByUsername(ctx context.Context, username string) (*model.UserEntity, error) {
	arguments := r.Mock.Called(ctx, username)

	var (
		return1 *model.UserEntity
		return2 error
	)

	if arguments.Get(0) != nil {
		return1 = arguments.Get(0).(*model.UserEntity)
	}
	if arguments.Get(1) != nil {
		return2 = arguments.Get(1).(error)
	}

	return return1, return2
}

func (r *UserRepositoryMock) GetUserByUsernamePassword(ctx context.Context, username string, password string) (*model.UserEntity, error) {
	arguments := r.Mock.Called(ctx, username, password)

	var (
		return1 *model.UserEntity
		return2 error
	)

	if arguments.Get(0) != nil {
		return1 = arguments.Get(0).(*model.UserEntity)
	}
	if arguments.Get(1) != nil {
		return2 = arguments.Get(1).(error)
	}

	return return1, return2
}

func (r *UserRepositoryMock) UpdateUser(ctx context.Context, userID int, updateUser *model.UserEntity) (*model.UserEntity, error) {
	arguments := r.Mock.Called(ctx, userID, updateUser)

	var (
		return1 *model.UserEntity
		return2 error
	)

	if arguments.Get(0) != nil {
		return1 = arguments.Get(0).(*model.UserEntity)
	}
	if arguments.Get(1) != nil {
		return2 = arguments.Get(1).(error)
	}

	return return1, return2
}

func (r *UserRepositoryMock) DeleteUser(ctx context.Context, userID int) error {
	arguments := r.Mock.Called(ctx, userID)

	var (
		return1 error
	)

	if arguments.Get(0) != nil {
		return1 = arguments.Get(0).(error)
	}

	return return1
}
