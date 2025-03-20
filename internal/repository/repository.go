package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository    IUserRepository
	ProductRepository IProductRepository
	StoreRepository   IStoreRepository
	CartRepository    ICartRepository
	AddressRepository IAddressRepository
	OrderRepository   IOrderRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:    NewUserRepository(db),
		ProductRepository: NewProductRepository(db),
		StoreRepository:   NewStoreRepository(db),
		CartRepository:    NewCartRepository(db),
		AddressRepository: NewAddressRepository(db),
		OrderRepository:   NewOrderRepository(db),
	}
}
