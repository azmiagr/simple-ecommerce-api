package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID    int         `json:"order_id" gorm:"type:int;primaryKey;autoIncrement"`
	TotalPrice float64     `json:"total_price" gorm:"type:decimal;not null"`
	Status     int         `json:"status" gorm:"type:int;default:0"` // 0 = pending, 1 = paid, 2 = canceled
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	UserID     uuid.UUID   `json:"user_id"`
	AddressID  int         `json:"address_id"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}
