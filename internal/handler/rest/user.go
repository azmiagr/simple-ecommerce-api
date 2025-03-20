package rest

import (
	"golang-ecommerce/model"
	"golang-ecommerce/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.UserService.Register(&param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register user", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "success to register new user", nil)
}

func (r *Rest) Login(ctx *gin.Context) {
	param := model.UserLogin{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	result, err := r.service.UserService.Login(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to login user", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success to login user", result)

}
