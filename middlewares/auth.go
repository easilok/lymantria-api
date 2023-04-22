package middlewares

import (
	"net/http"

	"github.com/easilok/lymantria-api/helpers"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		au, err := helpers.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("userId", au.UserId)
		c.Next()
	}
}
