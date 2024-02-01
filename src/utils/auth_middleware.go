package utils

import (
	"net/http"
	"strings"

	"auth/src/infrastructure/token"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message CodeError `json:"message"`
	Detail  string    `json:"detail"`
}

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := token.IsAuthorized(authToken, secret)
			if authorized {
				userId, err := token.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, ErrorResponse{Message: utils.LOGIC_CRASH, Detail: err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userId)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, ErrorResponse{Message: utils.UNAUTHORIZED, Detail: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: utils.UNAUTHORIZED, Detail: "Not authorized"})
		c.Abort()
	}
}