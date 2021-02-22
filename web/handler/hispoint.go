package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//HisPointGet Get his events
func HisPointGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/hispoint", gin.H{
		"systemTime":    str.NowString(),
		"hispointviews": service.GetHisPoint(),
	})
}
