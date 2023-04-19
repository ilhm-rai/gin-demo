package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/go-middleware/helper"
	"github.com/ilhm-rai/go-middleware/model"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helper.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.WebResponseMessage{
				Code: http.StatusUnauthorized,
				Status: "Unauthorized",
				Message: "Sign in to proceed",
			})
			return
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
