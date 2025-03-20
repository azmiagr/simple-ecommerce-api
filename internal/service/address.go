package service

import (
	"golang-ecommerce/entity"
	"golang-ecommerce/internal/repository"
	"golang-ecommerce/model"
)

type IAddressService interface {
	CreateAddress(param *model.CreateAddress) (*entity.Address, error)
	UpdateAddress(addressID int, userID string, param *model.UpdateAddress) (*entity.Address, error)
	DeleteAddress(addressID int, userID string) error
	GetAddress(userID string) ([]*entity.Address, error)
}

type AddressService struct {
	AddressRepository repository.IAddressRepository
}

func NewAddressService(AddressRepository repository.IAddressRepository) IAddressService {
	return &AddressService{
		AddressRepository: AddressRepository,
	}
}

func (as *AddressService) CreateAddress(param *model.CreateAddress) (*entity.Address, error) {
	address := &entity.Address{
		RecipentName:    param.RecipentName,
		PhoneNumber:     param.PhoneNumber,
		RecipentAddress: param.RecipentAddress,
		PostalCode:      param.PostalCode,
		UserID:          param.UserID,
	}

	address, err := as.AddressRepository.CreateAddress(address)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (as *AddressService) UpdateAddress(addressID int, userID string, param *model.UpdateAddress) (*entity.Address, error) {
	address, err := as.AddressRepository.UpdateAddress(addressID, userID, param)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (as *AddressService) DeleteAddress(addressID int, userID string) error {
	err := as.AddressRepository.DeleteAddress(addressID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (as *AddressService) GetAddress(userID string) ([]*entity.Address, error) {
	address, err := as.AddressRepository.GetAddress(userID)
	if err != nil {
		return nil, err
	}

	return address, nil
}
