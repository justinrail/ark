package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//HisLiveEventGet Get live events
func HisLiveEventGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/hisliveevent", gin.H{
		"systemTime":     str.NowString(),
		"liveeventviews": service.GetHisLiveEvent(),
	})
}
