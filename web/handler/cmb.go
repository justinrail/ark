package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CMBGet Get cmb collector data to page
func CMBGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/cmb", gin.H{
		"systemTime": str.NowString(),
		"metrics":    service.GetMetrics(),
	})

}
