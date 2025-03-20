package repository

import (
	"golang-ecommerce/entity"
	"golang-ecommerce/model"

	"gorm.io/gorm"
)

type IAddressRepository interface {
	CreateAddress(address *entity.Address) (*entity.Address, error)
	UpdateAddress(addressID int, userID string, param *model.UpdateAddress) (*entity.Address, error)
	DeleteAddress(addressID int, userID string) error
	GetAddress(userID string) ([]*entity.Address, error)
}

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) IAddressRepository {
	return &AddressRepository{db}
}

func (ar *AddressRepository) CreateAddress(address *entity.Address) (*entity.Address, error) {
	err := ar.db.Debug().Create(&address).Error
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (ar *AddressRepository) UpdateAddress(addressID int, userID string, param *model.UpdateAddress) (*entity.Address, error) {
	tx := ar.db.Begin()

	var address entity.Address
	err := tx.Debug().Where("address_id = ? && user_id = ?", addressID, userID).First(&address).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	addressParse := *parseUpdateAddress(param, &address)

	err = tx.Debug().Save(&addressParse).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &address, nil
}

func (ar *AddressRepository) DeleteAddress(addressID int, userID string) error {
	tx := ar.db.Begin()

	err := tx.Debug().Where("address_id = ? AND user_id = ?", addressID, userID).First(&entity.Address{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Debug().Where("address_id = ? AND user_id = ?", addressID, userID).Delete(&entity.Address{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (ar *AddressRepository) GetAddress(userID string) ([]*entity.Address, error) {
	var addresses []*entity.Address

	err := ar.db.Debug().Where("user_id = ?", userID).Find(&addresses).Error
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func parseUpdateAddress(model *model.UpdateAddress, address *entity.Address) *entity.Address {
	if model.RecipentName != "" {
		address.RecipentName = model.RecipentName
	}

	if model.PhoneNumber != "" {
		address.PhoneNumber = model.PhoneNumber
	}

	if model.RecipentAddress != "" {
		address.RecipentAddress = model.RecipentAddress
	}

	if model.PostalCode != "" {
		address.PostalCode = model.PostalCode
	}

	return address
}
