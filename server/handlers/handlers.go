package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzw200015/ServerStatus/server/cache/alive"
	"github.com/kzw200015/ServerStatus/server/cache/status"
	"github.com/kzw200015/ServerStatus/server/config"
	"github.com/kzw200015/ServerStatus/types"
)

type Node struct {
	Id      string `json:"id,omitempty"`
	IsAlive bool   `json:"is_alive"`
}

func HandleGetTitle(c *gin.Context) {
	c.String(http.StatusOK, config.Get().Title)
}

func HandleGetNodes(c *gin.Context) {
	var resp = []Node{}
	for _, node := range config.Get().Nodes {
		resp = append(resp, Node{Id: node.Id, IsAlive: alive.IsAlive(node.Id)})
	}

	c.JSON(http.StatusOK, resp)
}

func HandleGetNodeStatus(c *gin.Context) {
	id := c.Param("id")
	s, err := status.GetCache(id)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, s)
}

func HandlePostNodeStatus(c *gin.Context) {
	var s types.Status
	err := c.ShouldBind(&s)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	node := c.MustGet(gin.AuthUserKey).(string)
	status.SetCache(node, s)
}
