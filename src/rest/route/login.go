package route

import (
	"net/http"

	"auth/src/presentation/controller"
	"auth/src/infrastructure/mongo"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func LoginRouter(env *bootstrap.Env, time time.Duration, group *gin.RouterGroup) {
	ur := UserRepository(
		repository
	)

	group.GET("/users", Login)
}

func Login(c *gin.Context) {
	var request LoginRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	lc := controller.NewLogin(request.Email, request.Password)	

	c.JSON(http.StatusOK, lc)
}