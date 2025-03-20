package repository

import (
	"golang-ecommerce/entity"
	"golang-ecommerce/model"

	"gorm.io/gorm"
)

type ICartRepository interface {
	CreateCart(cart *entity.Cart) error
	GetCartByUserID(userID string) (*entity.Cart, error)
	AddItemToCart(cartItem *entity.CartItem) (*entity.CartItem, error)
	UpdateCartItemQuantity(param *model.AddToCart) (*entity.CartItem, error)
	GetCartItem(param *model.AddToCart) (*entity.CartItem, error)
	GetUserCartItemList(userID string) ([]*entity.CartItem, error)
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{db}
}

func (cr *CartRepository) CreateCart(cart *entity.Cart) error {
	err := cr.db.Debug().Create(&cart).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *CartRepository) GetCartByUserID(userID string) (*entity.Cart, error) {
	var cart entity.Cart
	err := cr.db.Debug().Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (cr *CartRepository) AddItemToCart(cartItem *entity.CartItem) (*entity.CartItem, error) {
	err := cr.db.Debug().Create(&cartItem).Error
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (cr *CartRepository) UpdateCartItemQuantity(param *model.AddToCart) (*entity.CartItem, error) {
	tx := cr.db.Begin()
	var cartItem entity.CartItem

	if param.Quantity > 0 {
		err := tx.Debug().Where("cart_id = ? && product_id = ?", param.CartID, param.ProductID).First(&cartItem).Update("quantity", param.Quantity).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		return &cartItem, nil
	}

	err := tx.Debug().Where("cart_id = ? AND product_id = ?", param.CartID, param.ProductID).First(&cartItem).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Debug().Where("cart_id = ? AND product_id = ?", param.CartID, param.ProductID).Delete(&cartItem).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return nil, nil

}

func (cr *CartRepository) GetCartItem(param *model.AddToCart) (*entity.CartItem, error) {
	var cartItem entity.CartItem
	err := cr.db.Debug().Where("cart_id = ? && product_id = ?", param.CartID, param.ProductID).First(&cartItem).Error
	if err != nil {
		return nil, err
	}

	return &cartItem, nil
}

func (cr *CartRepository) GetUserCartItemList(userID string) ([]*entity.CartItem, error) {
	var cartItems []*entity.CartItem

	err := cr.db.Debug().Preload("Product").Joins("JOIN carts ON cart_items.cart_id = carts.cart_id").Joins("JOIN products ON cart_items.product_id = products.product_id").Where("carts.user_id = ?", userID).Select("cart_items.*, products.product_name, products.price").Find(&cartItems).Error
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}
