package entity

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	StoreID          int       `json:"store_id" gorm:"type:int;primaryKey;autoIncrement"`
	StoreName        string    `json:"store_name" gorm:"type:varchar(100)"`
	StoreDescription string    `json:"store_description" gorm:"type:text"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserID           uuid.UUID `json:"user_id"`
	Products         []Product `json:"products" gorm:"foreignKey:StoreID"`
}
