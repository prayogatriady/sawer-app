package repository

import (
	"context"
	"time"

	"github.com/prayogatriady/sawer-app/model"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CreateUser(ctx context.Context, user *model.UserEntity) (*model.UserEntity, error)
	GetUser(ctx context.Context, userID int) (*model.UserEntity, error)
	GetUserByUsername(ctx context.Context, username string) (*model.UserEntity, error)
	GetUserByUsernamePassword(ctx context.Context, username string, password string) (*model.UserEntity, error)
	UpdateUser(ctx context.Context, userID int, updateUser *model.UserEntity) (*model.UserEntity, error)
	DeleteUser(ctx context.Context, userID int) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepoInterface {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, userEntity *model.UserEntity) (*model.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := ur.DB.WithContext(ctx).Table("users").Create(&userEntity).Error; err != nil {
		return userEntity, err
	}
	return userEntity, nil
}

func (ur *UserRepository) GetUser(ctx context.Context, userID int) (*model.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var user *model.UserEntity
	if err := ur.DB.WithContext(ctx).Table("users").Where("id =?", userID).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByUsername(ctx context.Context, username string) (*model.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var user *model.UserEntity
	if err := ur.DB.WithContext(ctx).Table("users").Where("username =?", username).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByUsernamePassword(ctx context.Context, username string, password string) (*model.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var user *model.UserEntity
	if err := ur.DB.WithContext(ctx).Table("users").Where("username =? AND password =?", username, password).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
func (ur *UserRepository) UpdateUser(ctx context.Context, userID int, updateUser *model.UserEntity) (*model.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var user *model.UserEntity
	if err := ur.DB.WithContext(ctx).Table("users").Where("id =?", userID).Updates(&updateUser).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
func (ur *UserRepository) DeleteUser(ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := ur.DB.WithContext(ctx).Table("users").Where("id =?", userID).Delete(&model.UserEntity{}).Error; err != nil {
		return err
	}
	return nil
}
