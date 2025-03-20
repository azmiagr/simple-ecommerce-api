package model

type AddToCart struct {
	CartID    int `json:"cart_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ReduceQuantityFromCart struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ViewCartResponse struct {
	CartItemID  int     `json:"cart_item_id"`
	ProductID   int     `json:"product_id"`
	CartID      int     `json:"cart_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
