package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/go-middleware/controllers"
)

func Start() *gin.Engine {
	engine := gin.Default()

	router := engine.Group("/users")
	{
		router.POST("/register", controllers.UserRegister)
	}

	return engine
}
