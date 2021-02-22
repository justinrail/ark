package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TopologyGet Get machine data to page
func TopologyGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/topology", gin.H{
		"systemTime": str.NowString(),
		"metrics":    service.GetMetrics(),
	})

}
