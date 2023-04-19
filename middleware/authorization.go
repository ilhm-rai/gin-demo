package middleware

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/go-middleware/model"
	"github.com/ilhm-rai/go-middleware/service"
)

func ProductAuthorization(service service.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.WebResponseMessage{
				Code:    http.StatusBadRequest,
				Status:  "Bad Request",
				Message: "Invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		role := userData["role"].(string)
		product, _ := service.FindProduct(uint(productId))

		if product == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, model.WebResponseMessage{
				Code:    http.StatusNotFound,
				Status:  "Not Found",
				Message: "Product not found",
			})
			return
		}

		if role == "USER" && product.UserID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.WebResponseMessage{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "You are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
