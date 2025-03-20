package model

type RegisterStore struct {
	StoreName   string `json:"store_name" binding:"required"`
	Description string `json:"description"`
}
