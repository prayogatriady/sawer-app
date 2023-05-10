package service

import (
	"context"
	"fmt"

	"github.com/prayogatriady/sawer-app/middleware"
	"github.com/prayogatriady/sawer-app/model"
	"github.com/prayogatriady/sawer-app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServInterface interface {
	Signup(ctx context.Context, userRequest *model.UserSignupRequest) (*model.UserResponse, error)
	Signin(ctx context.Context, userRequest *model.UserSigninRequest) (*model.UserTokenResponse, error)
	Profile(ctx context.Context, userId int) (*model.UserResponse, error)
	EditProfile(ctx context.Context, userId int, userRequest *model.UserEditRequest) (*model.UserResponse, error)
	DeleteUser(ctx context.Context, userId int) error
}

type UserService struct {
	UserRepository repository.UserRepoInterface
}

func NewUserService(userRepository repository.UserRepoInterface) UserServInterface {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) Signup(ctx context.Context, userRequest *model.UserSignupRequest) (*model.UserResponse, error) {
	// hashing passsword
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return &model.UserResponse{}, err
	}

	userEntity := &model.UserEntity{
		Username: userRequest.Username,
		Password: string(bytePassword),
	}

	userEntity, err = us.UserRepository.CreateUser(ctx, userEntity)
	if err != nil {
		return &model.UserResponse{}, err
	}

	userResponse := &model.UserResponse{
		ID:        userEntity.ID,
		Username:  userEntity.Username,
		Password:  userEntity.Password,
		Balance:   userEntity.Balance,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}

	return userResponse, nil
}

func (us *UserService) Signin(ctx context.Context, userRequest *model.UserSigninRequest) (*model.UserTokenResponse, error) {
	userFound, err := us.UserRepository.GetUserByUsername(ctx, userRequest.Username)
	if err != nil {
		return &model.UserTokenResponse{}, err
	}

	// comparing hash password in db with request password
	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userRequest.Password)); err != nil {
		return &model.UserTokenResponse{}, err
	}

	userEntity, err := us.UserRepository.GetUserByUsernamePassword(ctx, userFound.Username, userFound.Password)
	if err != nil {
		return &model.UserTokenResponse{}, fmt.Errorf("wrong password")
	}

	// generate token
	token, err := middleware.GenerateToken(userEntity)
	if err != nil {
		return &model.UserTokenResponse{}, err
	}

	return &model.UserTokenResponse{Token: token}, nil
}

func (us *UserService) Profile(ctx context.Context, userId int) (*model.UserResponse, error) {
	userEntity, err := us.UserRepository.GetUser(ctx, userId)
	if err != nil {
		return &model.UserResponse{}, err
	}

	userResponse := &model.UserResponse{
		ID:        userEntity.ID,
		Username:  userEntity.Username,
		Password:  userEntity.Password,
		Balance:   userEntity.Balance,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}

	return userResponse, nil
}
func (us *UserService) EditProfile(ctx context.Context, userId int, userRequest *model.UserEditRequest) (*model.UserResponse, error) {
	// hashing passsword
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return &model.UserResponse{}, err
	}

	// get current balance
	userCurrentEntity, err := us.UserRepository.GetUser(ctx, userId)
	if err != nil {
		return &model.UserResponse{}, err
	}

	userEntity := &model.UserEntity{
		ID:       userId,
		Username: userRequest.Username,
		Password: string(bytePassword),
		Balance:  userCurrentEntity.Balance + userRequest.Balance,
	}

	userEntity, err = us.UserRepository.UpdateUser(ctx, userId, userEntity)
	if err != nil {
		return &model.UserResponse{}, err
	}

	userResponse := &model.UserResponse{
		ID:        userEntity.ID,
		Username:  userEntity.Username,
		Password:  userEntity.Password,
		Balance:   userEntity.Balance,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}

	return userResponse, nil
}
func (us *UserService) DeleteUser(ctx context.Context, userId int) error {
	if err := us.UserRepository.DeleteUser(ctx, userId); err != nil {
		return err
	}
	return nil
}
