package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kzw200015/ServerStatus/assets"
	"github.com/kzw200015/ServerStatus/server/config"
	"github.com/kzw200015/ServerStatus/server/handlers"
	"github.com/kzw200015/ServerStatus/server/middlewares"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.HandleErrors()).Use(middlewares.HandleStatics(assets.EFS))
	api := r.Group("/api").Use(cors.Default())
	{
		api.GET("/title", handlers.HandleGetTitle)
		api.GET("/nodes", handlers.HandleGetNodes)
		api.GET("/nodes/:id/status", handlers.HandleGetNodeStatus)
		api.POST("/nodes", gin.BasicAuth(config.GetAuthMap()), middlewares.HandleAlive(), handlers.HandlePostNodeStatus)
	}

	return r
}
