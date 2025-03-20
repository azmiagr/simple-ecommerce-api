package entity

import "time"

type OrderItem struct {
	OrderItemID int       `json:"order_item_id" gorm:"type:int;primaryKey;autoIncrement"`
	Quantity    int       `json:"quantity" gorm:"type:int;not null"`
	Price       float64   `json:"price" gorm:"type:decimal;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	OrderID     int       `json:"order_id"`
	ProductID   int       `json:"product_id"`
}
