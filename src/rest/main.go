package main

import (
	"time"

	"auth/src/infrastructure/bootstrap"
	"auth/src/rest/route"

	"github.com/gin-gonic/gin"
)
func main() {
	app := bootstrap.App()
	env := app.Env
	timeout := time.Duration(env.ContextTimeout) * time.Second

	r := gin.Default()
	route.Setup(env, timeout, r)
	r.Run(env.ServerAddress)
}