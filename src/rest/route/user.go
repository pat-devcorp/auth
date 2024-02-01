package route

import (
	"time"

	"auth/src/domain"
	"auth/src/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func UserRouter(env *bootstrap.Env, time time.Duration, group *gin.RouterGroup) {
	// ur := UserRepository(db, domain.CollectionUser)
	// lc := &controller.LoginController{
	// 	LoginApplication: application.NewLogin(ur, timeout),
	// 	Env: env
	// }

	group.GET("/users", func(c *gin.Context) {
		c.IndentedJSON(200, domain.Users)
	})
}