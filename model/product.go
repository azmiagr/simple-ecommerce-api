package model

type GetAllProducts struct {
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	StoreName   string  `json:"store_name"`
}

type SearchProduct struct {
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	StoreName   string  `json:"store_name"`
}

type AddProductRequest struct {
	ProductName        string  `json:"product_name" binding:"required"`
	ProductDescription string  `json:"product_description"`
	Price              float64 `json:"price" binding:"required"`
	StoreID            int     `json:"store_id"`
	Stock              int     `json:"stock" binding:"required"`
}
