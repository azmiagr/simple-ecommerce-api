package rest

import (
	"golang-ecommerce/entity"
	"golang-ecommerce/model"
	"golang-ecommerce/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateProduct(ctx *gin.Context) {
	var param model.AddProductRequest
	user := ctx.MustGet("user").(*entity.User)
	param.StoreID = user.Store.StoreID

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid request", err)
		return
	}

	product, err := r.service.ProductService.CreateProduct(&param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to add product", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success to add product", product)

}

func (r *Rest) GetAllProducts(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid page", err)
		return
	}

	products, err := r.service.ProductService.GetAllProducts(page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get all products", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success to get product", products)
}

func (r *Rest) GetProductByName(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid page", err)
		return
	}

	productName := ctx.Query("productName")
	product, err := r.service.ProductService.GetProductByName(productName, page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get product", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success to get product", product)
}
