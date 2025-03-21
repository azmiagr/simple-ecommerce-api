package repository

import (
	"golang-ecommerce/entity"
	"golang-ecommerce/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUser(param model.UserParam) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	err := ur.db.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUser(param model.UserParam) (*entity.User, error) {
	user := entity.User{}
	err := ur.db.Debug().Preload("Store").Preload("Address").Preload("Cart").Preload("Orders").Where(&param).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
