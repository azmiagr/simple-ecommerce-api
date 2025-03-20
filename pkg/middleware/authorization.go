package middleware

import (
	"errors"
	"golang-ecommerce/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *middleware) OnlyRegisterUser(ctx *gin.Context) {
	user, err := m.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		response.Error(ctx, http.StatusForbidden, "failed get login user", err)
		ctx.Abort()
		return
	}

	if user.RoleID != 2 {
		response.Error(ctx, http.StatusForbidden, "this endpoint cannot be access", errors.New("user dont have access"))
		ctx.Abort()
		return
	}
	ctx.Next()
}
