package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/kzw200015/ServerStatus/server/cache/alive"
)

func HandleAlive() gin.HandlerFunc {
	return func(c *gin.Context) {
		alive.Refresh(c.MustGet(gin.AuthUserKey).(string))
	}
}
