package middlewares

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleStatics(efs embed.FS) gin.HandlerFunc {
	sFS, err := fs.Sub(efs, "frontend/dist")
	if err != nil {
		log.Panicln(err)
	}
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// API 跳过
		if strings.HasPrefix(path, "/api") {
			c.Next()
			return
		}

		//此处处理静态资源
		fileServer := http.FileServer(http.FS(sFS))
		fileServer.ServeHTTP(c.Writer, c.Request)
		//处理完静态资源，调用c.Abort()中断请求
		c.Abort()
	}
}
