package route

import (
	"net/http"

	"auth/src/presentation/controller"
	"auth/src/infrastructure/mongo/repository"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func LoginRouter(time time.Duration, group *gin.RouterGroup) {
	ur := repository.SetDefault()

	group.GET("/users", Login)
}

func Login(c *gin.Context) {
	var request LoginRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	lc, err := controller.NewLogin(request.Email, request.Password)	
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, lc)
}