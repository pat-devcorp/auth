package route

import (
	"time"

	"auth/src/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func HelloRouter(env *bootstrap.Env, time time.Duration, group *gin.RouterGroup) {
	// ur := UserRepository(db, domain.CollectionUser)
	// lc := &controller.LoginController{
	// 	LoginApplication: application.NewLogin(ur, timeout),
	// 	Env: env
	// }

	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}