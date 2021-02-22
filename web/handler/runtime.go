package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RuntimeGet Get Runtime data to page
func RuntimeGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/runtime", gin.H{
		"systemTime": str.NowString(),
		"runtimes":   service.GetAllRuntime(),
	})

}
