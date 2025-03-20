package model

import "golang-ecommerce/entity"

type CheckoutRequest struct {
	AddressID  int   `json:"address_id"`
	CartItemID []int `json:"cart_item_ids"`
}

type CreateOrder struct {
	UserID     string             `json:"user_id"`
	AddressID  int                `json:"address_id"`
	TotalPrice float64            `json:"total_price"`
	Status     int                `json:"status"`
	OrderItems []entity.OrderItem `json:"order_items"`
}
