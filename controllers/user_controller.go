package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/go-middleware/database"
	"github.com/ilhm-rai/go-middleware/helpers"
	"github.com/ilhm-rai/go-middleware/models"
)

var (
	jsonContentType = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.User{}

	if contentType == jsonContentType {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.FullName,
		"role":      User.Role,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}
	password := ""

	if contentType == jsonContentType {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	isPasswordValid := helpers.ComparePass([]byte(User.Password), []byte(password))

	if err != nil || !isPasswordValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, string(User.Role))

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
