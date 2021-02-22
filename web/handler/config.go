package handler

import (
	"net/http"

	"ark/util/str"
	"ark/web/service"

	"github.com/gin-gonic/gin"
)

//ConfigGet Web Path /config handler
func ConfigGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/config", gin.H{
		"systemTime": str.NowString(),
		"cfgs":       service.GetAllConfig(),
	})

}
