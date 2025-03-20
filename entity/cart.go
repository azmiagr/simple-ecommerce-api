package entity

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	CartID    int        `json:"cart_id" gorm:"type:int;primaryKey;autoIncrement"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	UserID    uuid.UUID  `json:"user_id"`
	CartItems []CartItem `gorm:"foreignKey:CartID"`
}
