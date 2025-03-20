package rest

import (
	"golang-ecommerce/entity"
	"golang-ecommerce/model"
	"golang-ecommerce/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateAddress(ctx *gin.Context) {
	var param model.CreateAddress
	user := ctx.MustGet("user").(*entity.User)
	param.UserID = user.UserID
	err := ctx.ShouldBind(&param)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind input", err)
		return
	}

	address, err := r.service.AddressService.CreateAddress(&param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to create address", err)
		return
	}

	addressResponse := model.CreateAddress{
		UserID:          address.UserID,
		RecipentName:    address.RecipentName,
		PhoneNumber:     address.PhoneNumber,
		RecipentAddress: address.RecipentAddress,
		PostalCode:      address.PostalCode,
	}

	response.Success(ctx, http.StatusOK, "success create address", addressResponse)

}

func (r *Rest) UpdateAddress(ctx *gin.Context) {
	idParam := ctx.Param("address_id")
	var param model.UpdateAddress

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to bind input", err)
		return
	}

	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid id", err)
		return
	}

	user := ctx.MustGet("user").(*entity.User)
	userID := user.UserID.String()

	address, err := r.service.AddressService.UpdateAddress(idInt, userID, &param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to update address", err)
		return
	}

	addressResponse := model.UpdateAddress{
		RecipentName:    address.RecipentName,
		PhoneNumber:     address.PhoneNumber,
		RecipentAddress: address.RecipentAddress,
		PostalCode:      address.PostalCode,
	}

	response.Success(ctx, http.StatusOK, "success update address", addressResponse)
}

func (r *Rest) DeleteAddress(ctx *gin.Context) {
	idParam := ctx.Param("address_id")
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid address id", err)
		return
	}

	user := ctx.MustGet("user").(*entity.User)
	userID := user.UserID.String()

	err = r.service.AddressService.DeleteAddress(idInt, userID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to delete address", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success to delete address", nil)
}

func (r *Rest) GetAddress(ctx *gin.Context) {
	user := ctx.MustGet("user").(*entity.User)
	userID := user.UserID.String()

	addresses, err := r.service.AddressService.GetAddress(userID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get addresses", err)
		return
	}

	var addressResponse []model.AddressResponse
	for _, a := range addresses {
		addressResponse = append(addressResponse, model.AddressResponse{
			RecipentName:    a.RecipentName,
			PhoneNumber:     a.PhoneNumber,
			RecipentAddress: a.RecipentAddress,
			PostalCode:      a.PostalCode,
			UserID:          userID,
		})
	}

	response.Success(ctx, http.StatusOK, "success get addresses", addressResponse)
}
