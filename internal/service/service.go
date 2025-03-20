package service

import (
	"golang-ecommerce/internal/repository"
	"golang-ecommerce/pkg/bcrypt"
	"golang-ecommerce/pkg/jwt"
)

type Service struct {
	UserService    IUserService
	ProductService IProductService
	StoreService   IStoreService
	CartService    ICartService
	AddressService IAddressService
	OrderService   IOrderService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) *Service {
	return &Service{
		UserService:    NewUserService(repository.UserRepository, bcrypt, jwtAuth, repository.CartRepository),
		ProductService: NewProductService(repository.ProductRepository),
		StoreService:   NewStoreService(repository.StoreRepository),
		CartService:    NewCartService(repository.CartRepository),
		AddressService: NewAddressService(repository.AddressRepository),
		OrderService:   NewOrderService(repository.OrderRepository),
	}
}
