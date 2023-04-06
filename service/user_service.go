package service

import (
	"context"

	"github.com/prayogatriady/sawer-app/model"
	"github.com/prayogatriady/sawer-app/repository"
)

type UserServInterface interface {
	CreateUser(ctx context.Context, userRequest *model.UserSignupRequest) (*model.UserResponse, error)
}

type UserService struct {
	UserRepository repository.UserRepoInterface
}

func NewUserService(userRepository repository.UserRepoInterface) UserServInterface {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) CreateUser(ctx context.Context, userRequest *model.UserSignupRequest) (*model.UserResponse, error) {
	var userEntity *model.UserEntity
	userEntity = &model.UserEntity{
		Username: userRequest.Username,
		Password: userRequest.Password,
	}

	userEntity, err := us.UserRepository.CreateUser(ctx, userEntity)
	if err != nil {
		return &model.UserResponse{}, err
	}

	var userResponse *model.UserResponse
	userResponse = &model.UserResponse{
		ID:        userEntity.ID,
		Username:  userEntity.Username,
		Password:  userEntity.Password,
		Balance:   userEntity.Balance,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}

	return userResponse, nil
}
