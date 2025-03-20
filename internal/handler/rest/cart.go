package rest

import (
	"golang-ecommerce/entity"
	"golang-ecommerce/model"
	"golang-ecommerce/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) AddToCart(ctx *gin.Context) {
	var param model.AddToCart
	user := ctx.MustGet("user").(*entity.User)
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind input", err)
		return
	}

	userID := user.UserID.String()
	param.CartID = user.Cart.CartID
	cartItem, err := r.service.CartService.AddToCart(userID, &param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to add cart item", err)
		return
	}

	cartResponse := model.AddToCart{
		CartID:    user.Cart.CartID,
		ProductID: cartItem.ProductID,
		Quantity:  cartItem.Quantity,
	}
	response.Success(ctx, http.StatusOK, "success add item to cart", cartResponse)

}

func (r *Rest) RemoveFromCart(ctx *gin.Context) {
	var param model.AddToCart
	user := ctx.MustGet("user").(*entity.User)
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind input", err)
		return
	}

	userID := user.UserID.String()
	cartItem, err := r.service.CartService.RemoveFromCart(userID, &param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to edit cart items", err)
		return
	}

	if cartItem == nil {
		response.Success(ctx, http.StatusOK, "item successfully deleted", nil)
	} else {
		cartResponse := model.ReduceQuantityFromCart{
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
		}
		response.Success(ctx, http.StatusOK, "success reduces item from cart", cartResponse)
	}

}

func (r *Rest) GetUserCartItemList(ctx *gin.Context) {
	user := ctx.MustGet("user").(*entity.User)
	userID := user.UserID.String()

	cartItems, err := r.service.CartService.GetUserCartItemList(userID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get cart items", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success get cart items", cartItems)
}
