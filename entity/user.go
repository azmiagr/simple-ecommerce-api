package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"user_id" gorm:"type:varchar(36);primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(50);not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	RoleID    int       `json:"role_id"`
	Cart      Cart      `json:"cart" gorm:"foreignKey:UserID"`
	Store     Store     `json:"store" gorm:"foreignKey:UserID"`
	Address   []Address `json:"address" gorm:"foreignKey:UserID"`
	Orders    []Order   `json:"orders" gorm:"foreignKey:UserID"`
}
