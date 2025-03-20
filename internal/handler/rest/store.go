package rest

import (
	"golang-ecommerce/entity"
	"golang-ecommerce/model"
	"golang-ecommerce/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) RegisterStore(ctx *gin.Context) {
	var param model.RegisterStore
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind json input", err)
		return
	}

	user := ctx.MustGet("user").(*entity.User)
	store, err := r.service.StoreService.RegisterStore(user.UserID.String(), param.StoreName, param.Description)

	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register store", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success create new store", store)
}
