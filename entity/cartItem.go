package entity

import "time"

type CartItem struct {
	CartItemID int       `json:"cart_item_id" gorm:"type:int;primaryKey;autoIncrement"`
	Quantity   int       `json:"quantity" gorm:"type:int;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CartID     int       `json:"cart_id"`
	ProductID  int       `json:"product_id"`
	Product    Product
}
