package rest

import (
	"fmt"
	"golang-ecommerce/internal/service"
	"golang-ecommerce/pkg/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {

	routerGroup := r.router.Group("/api/v1")
	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)
	routerGroup.GET("/get-products", r.GetAllProducts)
	routerGroup.GET("/search-product", r.GetProductByName)

	user := routerGroup.Group("/users")
	user.Use(r.middleware.AuthenticateUser, r.middleware.OnlyRegisterUser)
	user.POST("/register-store", r.RegisterStore)
	user.POST("/add-product", r.CreateProduct)
	user.POST("/add-to-cart", r.AddToCart)
	user.POST("/remove-from-cart", r.RemoveFromCart)
	user.POST("/add-address", r.CreateAddress)
	user.POST("/checkout", r.Checkout)
	user.PATCH("/update-address/:address_id", r.UpdateAddress)
	user.GET("/view-cart", r.GetUserCartItemList)
	user.GET("/get-address", r.GetAddress)
	user.DELETE("/delete-address/:address_id", r.DeleteAddress)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
