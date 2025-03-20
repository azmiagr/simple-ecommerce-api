package rest

import (
	"golang-ecommerce/entity"
	"golang-ecommerce/model"
	"golang-ecommerce/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Checkout(ctx *gin.Context) {
	var req model.CheckoutRequest
	user := ctx.MustGet("user").(*entity.User)
	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	order, err := r.service.OrderService.Checkout(user.UserID, user.Cart.CartID, &req)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed checkout order", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success checkout order", order)
}
