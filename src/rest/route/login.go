package route

import (
	"net/http"

	"auth/src/infrastructure/bootstrap"
	"auth/src/presentation/controller"
	"auth/src/infrastructure/mongo/repository"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func LoginRouter(env *bootstrap.Env, time time.Duration, group *gin.RouterGroup) {
	group.GET("/user", login)
}

func login(c *gin.Context) {
	var request LoginRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	lc, err := &controller.Login(request.Email, request.Password)	
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, lc)
}