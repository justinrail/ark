package handler

import (
	"net/http"

	"ark/util/str"
	"ark/web/service"

	"github.com/gin-gonic/gin"
)

//LogGet Web Path /log Get handler
func LogGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/log", gin.H{
		"systemTime":  str.NowString(),
		"LogContents": service.GetAllLog(),
	})

}
