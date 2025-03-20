package entity

import "time"

type Product struct {
	ProductID          int         `json:"product_id" gorm:"type:int;primaryKey;autoIncrement"`
	ProductName        string      `json:"product_name" gorm:"type:varchar(100);not null"`
	ProductDescription string      `json:"product_description" gorm:"type:text;not null"`
	Price              float64     `json:"price" gorm:"type:decimal;not null"`
	Stock              int         `json:"stock" gorm:"type:int;not null"`
	ImageURL           string      `json:"image_url" gorm:"type:varchar(255)"`
	CreatedAt          time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	StoreID            int         `json:"store_id"`
	OrderItems         []OrderItem `gorm:"foreignKey:ProductID"`
	CartItems          []CartItem  `gorm:"foreignKey:ProductID"`
}
