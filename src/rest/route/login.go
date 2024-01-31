package route

import (
	"time"

	"auth/src/presentation/controller"
	"auth/src/infrastructure/mongo/user"
	"auth/src/domain"
	"auth/src/application"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *bootstrap.Env, time time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := UserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginApplication: application.NewLogin(ur, timeout),
		Env: env
	}
	group.POST("/login", lc.Login)
}
