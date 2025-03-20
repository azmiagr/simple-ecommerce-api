package mariadb

import (
	"golang-ecommerce/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
		&entity.Cart{},
		&entity.Address{},
		&entity.Store{},
		&entity.Product{},
		&entity.Order{},
		&entity.OrderItem{},
		&entity.CartItem{},
	)
	if err != nil {
		return err
	}

	return nil
}
