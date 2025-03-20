package repository

import (
	"golang-ecommerce/entity"

	"gorm.io/gorm"
)

type IOrderRepository interface {
	GetCartItemsByID(cartID int, itemIDs []int) ([]*entity.CartItem, error)
	CreateOrder(order *entity.Order) (*entity.Order, error)
	DeleteCartItemsByID(cartID int, itemIDs []int) error
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) IOrderRepository {
	return &OrderRepository{db}
}

func (or *OrderRepository) GetCartItemsByID(cartID int, itemIDs []int) ([]*entity.CartItem, error) {
	var cartItems []*entity.CartItem
	err := or.db.Debug().Preload("Product").Where("cart_id = ? AND cart_item_id IN ?", cartID, itemIDs).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (or *OrderRepository) CreateOrder(order *entity.Order) (*entity.Order, error) {
	err := or.db.Debug().Create(order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (or *OrderRepository) DeleteCartItemsByID(cartID int, itemIDs []int) error {
	err := or.db.Debug().Where("cart_id = ? AND cart_item_id IN ?", cartID, itemIDs).Delete(entity.CartItem{}).Error
	if err != nil {
		return err
	}

	return nil
}
