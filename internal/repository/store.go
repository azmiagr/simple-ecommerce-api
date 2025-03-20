package repository

import (
	"golang-ecommerce/entity"

	"gorm.io/gorm"
)

type IStoreRepository interface {
	CreateStore(store *entity.Store) error
	GetStoreByUserID(userID string) (*entity.Store, error)
}

type StoreRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) IStoreRepository {
	return &StoreRepository{db}
}

func (sr *StoreRepository) CreateStore(store *entity.Store) error {
	err := sr.db.Debug().Create(&store).Error
	if err != nil {
		return err
	}

	return nil
}

func (sr *StoreRepository) GetStoreByUserID(userID string) (*entity.Store, error) {
	var store entity.Store
	err := sr.db.Where("user_id = ?", userID).First(&store).Error
	if err != nil {
		return nil, err
	}

	return &store, nil
}
