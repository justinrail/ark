package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//StubGet Get stub metrics data to page
func StubGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/stub", gin.H{
		"systemTime": str.NowString(),
		"metrics":    service.GetMetrics(),
	})

}
