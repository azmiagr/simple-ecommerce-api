package service

import (
	"errors"
	"golang-ecommerce/entity"
	"golang-ecommerce/internal/repository"
	"golang-ecommerce/model"

	"github.com/google/uuid"
)

type IOrderService interface {
	Checkout(userID uuid.UUID, cartID int, param *model.CheckoutRequest) (*entity.Order, error)
}

type OrderService struct {
	orderRepo repository.IOrderRepository
}

func NewOrderService(orderRepo repository.IOrderRepository) IOrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (os *OrderService) Checkout(userID uuid.UUID, cartID int, param *model.CheckoutRequest) (*entity.Order, error) {
	cartItems, err := os.orderRepo.GetCartItemsByID(cartID, param.CartItemID)
	if err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, errors.New("no items selected for checkout")
	}

	var totalPrice float64
	var orderItems []entity.OrderItem

	for _, item := range cartItems {
		orderItems = append(orderItems, entity.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})

		totalPrice += item.Product.Price * float64(item.Quantity)
	}

	order := &entity.Order{
		UserID:     userID,
		AddressID:  param.AddressID,
		TotalPrice: totalPrice,
		Status:     0,
		OrderItems: orderItems,
	}

	order, err = os.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	err = os.orderRepo.DeleteCartItemsByID(cartID, param.CartItemID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
