package middleware

import (
	"golang-ecommerce/internal/service"
	"golang-ecommerce/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Interface interface {
	OnlyRegisterUser(ctx *gin.Context)
	AuthenticateUser(ctx *gin.Context)
}

type middleware struct {
	service *service.Service
	jwtAuth jwt.Interface
}

func Init(service *service.Service, jwtAuth jwt.Interface) Interface {
	return &middleware{
		service: service,
		jwtAuth: jwtAuth,
	}
}
