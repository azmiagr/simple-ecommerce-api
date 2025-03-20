package entity

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	AddressID       int       `json:"address_id" gorm:"type:int;primaryKey;autoIncrement"`
	RecipentName    string    `json:"recipent_name" gorm:"type:varchar(255);not null"`
	PhoneNumber     string    `json:"phone_number" gorm:"type:varchar(20);not null"`
	RecipentAddress string    `json:"recipent_address" gorm:"type:text;not null"`
	PostalCode      string    `json:"postal_code" gorm:"type:varchar(10);not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserID          uuid.UUID `json:"user_id"`
	Orders          []Order   `gorm:"foreignKey:AddressID"`
}
