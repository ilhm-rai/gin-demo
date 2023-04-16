package controller

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/go-middleware/exception"
	"github.com/ilhm-rai/go-middleware/middleware"
	"github.com/ilhm-rai/go-middleware/model"
	"github.com/ilhm-rai/go-middleware/service"
	"github.com/ilhm-rai/go-middleware/validation"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{ProductService: *productService}
}

func (controller *ProductController) Route(app *gin.Engine) {
	productRouter := app.Group("/products")
	{
		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", controller.CreateProduct)
		productRouter.GET("/", controller.FindProducts)
		productRouter.GET("/:productId", middleware.ProductAuthorization(controller.ProductService), controller.FindProduct)
		productRouter.PUT("/:productId", middleware.ProductAuthorization(controller.ProductService), controller.UpdateProduct)
		productRouter.DELETE("/:productId", middleware.ProductAuthorization(controller.ProductService), controller.DeleteProduct)
	}
}

func (controller *ProductController) CreateProduct(c *gin.Context) {
	var request model.CreateProductRequest
	err := c.ShouldBindJSON(&request)
	exception.PanicIfNeeded(err)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err = validation.Validate(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.WebResponseMessage{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	request.UserID = userID
	newProduct, err := controller.ProductService.CreateProduct(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.WebResponseMessage{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   newProduct,
	})
}

func (controller *ProductController) UpdateProduct(c *gin.Context) {
	var request model.UpdateProductRequest
	c.ShouldBindJSON(&request)

	err := validation.Validate(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponseMessage{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	request.UserID = userID

	err = controller.ProductService.UpdateProduct(uint(productId), request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.WebResponseMessage{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponseMessage{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Product successfully updated",
	})
}

func (controller *ProductController) FindProducts(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	role := userData["role"].(string)

	var err error
	var products []model.ProductResponse

	if role == "ADMIN" {
		products, err = controller.ProductService.FindProducts()
	} else {
		products, err = controller.ProductService.FindProductsByUserId(userID)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.WebResponseMessage{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   products,
	})
}

func (controller *ProductController) FindProduct(c *gin.Context) {
	var err error
	productId, err := strconv.Atoi(c.Param("productId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponseMessage{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	product, err := controller.ProductService.FindProduct(uint(productId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.WebResponseMessage{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	})
}

func (controller *ProductController) DeleteProduct(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))
	err := controller.ProductService.DeleteProduct(uint(productId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.WebResponseMessage{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
