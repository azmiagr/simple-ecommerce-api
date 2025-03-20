package middleware

import (
	"golang-ecommerce/model"
	"golang-ecommerce/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.Error(ctx, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	token := strings.Split(bearer, " ")[1]

	userID, err := m.jwtAuth.ValidateToken(token)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "unauthorized, ", err)
		ctx.Abort()
		return
	}

	user, err := m.service.UserService.GetUser(model.UserParam{
		UserID: userID,
	})
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "unauthorized, ", err)
		ctx.Abort()
		return
	}

	ctx.Set("user", user)
	ctx.Next()

}
