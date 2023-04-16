package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/go-middleware/config"
	"github.com/ilhm-rai/go-middleware/controller"
	"github.com/ilhm-rai/go-middleware/exception"
	"github.com/ilhm-rai/go-middleware/repository"
	"github.com/ilhm-rai/go-middleware/service"
)

func main() {
	configuration := config.New()
	database := config.NewPostgresDatabase(configuration)

	// Setup Repository
	userRepository := repository.NewUserRepository(database)
	productRepository := repository.NewProductRepository(database)

	// Setup Service
	userService := service.NewUserService(&userRepository)
	productService := service.NewProductService(&productRepository)

	// Setup Controller
	userController := controller.NewUserController(&userService)
	productController := controller.NewProductController(&productService)

	app := gin.Default()

	// Setup Routing
	userController.Route(app)
	productController.Route(app)

	// Start App
	err := app.Run(":8080")
	exception.PanicIfNeeded(err)
}
