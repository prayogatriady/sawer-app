package repository

import (
	"context"

	"github.com/prayogatriady/sawer-app/model"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CreateTables() error
	CreateUser(ctx context.Context, user *model.UserEntity) (*model.UserEntity, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepoInterface {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) CreateTables() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id BIGINT NOT NULL AUTO_INCREMENT,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(255) NOT NULL,
		balance INT DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP DEFAULT NULL,
		INDEX (id),
		PRIMARY KEY (id)
	)`

	if err := ur.DB.Exec(query).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, userEntity *model.UserEntity) (*model.UserEntity, error) {
	if err := ur.DB.Table("users").Create(&userEntity).Error; err != nil {
		return userEntity, err
	}
	return userEntity, nil
}
