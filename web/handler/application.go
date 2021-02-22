package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ApplicationGet Get app data to page
func ApplicationGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/application", gin.H{
		"systemTime": str.NowString(),
		"apps":       service.GetAllApplication(),
	})

}
