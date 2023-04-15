package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/go-middleware/controllers"
	"github.com/ilhm-rai/go-middleware/middlewares"
)

func Start() *gin.Engine {
	engine := gin.Default()

	userRouter := engine.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := engine.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.GET("/", controllers.FindProducts)
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productId", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return engine
}
