package service

import (
	"errors"
	"golang-ecommerce/entity"
	"golang-ecommerce/internal/repository"

	"github.com/google/uuid"
)

type IStoreService interface {
	RegisterStore(userID, storeName, description string) (*entity.Store, error)
}

type StoreService struct {
	StoreRepository repository.IStoreRepository
}

func NewStoreService(storeRepository repository.IStoreRepository) IStoreService {
	return &StoreService{
		StoreRepository: storeRepository,
	}
}

func (ss *StoreService) RegisterStore(userID, storeName, description string) (*entity.Store, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	existingStore, _ := ss.StoreRepository.GetStoreByUserID(userID)
	if existingStore != nil {
		return nil, errors.New("store already exists")
	}

	newStore := &entity.Store{
		StoreName:        storeName,
		StoreDescription: description,
		UserID:           userUUID,
	}

	err = ss.StoreRepository.CreateStore(newStore)
	if err != nil {
		return nil, err
	}

	return newStore, nil
}
