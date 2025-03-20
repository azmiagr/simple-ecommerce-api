package service

import (
	"errors"
	"golang-ecommerce/entity"
	"golang-ecommerce/internal/repository"
	"golang-ecommerce/model"
)

type ICartService interface {
	AddToCart(userID string, param *model.AddToCart) (*entity.CartItem, error)
	RemoveFromCart(userID string, param *model.AddToCart) (*entity.CartItem, error)
	GetUserCartItemList(userID string) ([]*model.ViewCartResponse, error)
}

type CartService struct {
	CartRepository repository.ICartRepository
}

func NewCartService(cartRepository repository.ICartRepository) ICartService {
	return &CartService{
		CartRepository: cartRepository,
	}
}

func (cs *CartService) AddToCart(userID string, param *model.AddToCart) (*entity.CartItem, error) {
	cart, err := cs.CartRepository.GetCartByUserID(userID)
	if err != nil {
		return nil, errors.New("cart not found for user")
	}

	cartItem, err := cs.CartRepository.GetCartItem(param)
	if err == nil {
		newQuantity := cartItem.Quantity + param.Quantity
		param.Quantity = newQuantity
		cartItemUpdate, err := cs.CartRepository.UpdateCartItemQuantity(param)

		if err != nil {
			return nil, err
		}

		return cartItemUpdate, nil
	}

	newCartItem := &entity.CartItem{
		CartID:    cart.CartID,
		ProductID: param.ProductID,
		Quantity:  param.Quantity,
	}

	item, err := cs.CartRepository.AddItemToCart(newCartItem)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (cs *CartService) RemoveFromCart(userID string, param *model.AddToCart) (*entity.CartItem, error) {
	cart, err := cs.CartRepository.GetCartByUserID(userID)
	if err != nil {
		return nil, errors.New("cart not found for user")
	}

	param.CartID = cart.CartID
	cartItem, err := cs.CartRepository.GetCartItem(param)
	if err != nil {
		return nil, errors.New("product not found in cart")
	}

	newQuantity := cartItem.Quantity - param.Quantity
	param.Quantity = newQuantity

	cartUpdate, err := cs.CartRepository.UpdateCartItemQuantity(param)
	if err != nil {
		return nil, err
	}

	return cartUpdate, nil
}

func (cs *CartService) GetUserCartItemList(userID string) ([]*model.ViewCartResponse, error) {
	items, err := cs.CartRepository.GetUserCartItemList(userID)
	if err != nil {
		return nil, err
	}

	var response []*model.ViewCartResponse
	for _, item := range items {
		response = append(response, &model.ViewCartResponse{
			CartItemID:  item.CartItemID,
			ProductID:   item.ProductID,
			CartID:      item.CartID,
			ProductName: item.Product.ProductName,
			Price:       item.Product.Price,
			Quantity:    item.Quantity,
		})
	}

	return response, nil
}
