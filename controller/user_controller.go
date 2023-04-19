package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/go-middleware/exception"
	"github.com/ilhm-rai/go-middleware/model"
	"github.com/ilhm-rai/go-middleware/service"
	"github.com/ilhm-rai/go-middleware/validation"
	"gorm.io/gorm"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{UserService: *userService}
}

func (controller *UserController) Route(app *gin.Engine) {
	userRouter := app.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
	}
}

func (controller *UserController) UserRegister(c *gin.Context) {
	var request model.RegisterUserRequest

	err := c.ShouldBindJSON(&request)

	exception.PanicIfNeeded(err)

	err = validation.Validate(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.WebResponseMessage{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	response, err := controller.UserService.Register(request)

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		c.AbortWithStatusJSON(http.StatusConflict, model.WebResponseMessage{
			Code:    http.StatusConflict,
			Status:  "Conflict",
			Message: "Email already registered",
		})
		return
	}

	c.JSON(http.StatusCreated, model.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *UserController) UserLogin(c *gin.Context) {
	var request model.LoginUserRequest
	err := c.ShouldBindJSON(&request)
	exception.PanicIfNeeded(err)

	token, err := controller.UserService.Login(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.WebResponseMessage{
			Code:    http.StatusUnauthorized,
			Status:  "Unauthorized",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: gin.H{
			"token": token,
		},
	})
}
