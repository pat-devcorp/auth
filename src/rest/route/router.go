package route

import (
	"time"

	"auth/src/infrastructure/bootstrap"
	"auth/src/utils"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("")
	HelloRouter(env, timeout, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(utils.JwtAuthMiddleware(env.AccessTokenSecret))
	UserRouter(env, timeout, publicRouter)
}