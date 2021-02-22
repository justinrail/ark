package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//LiveEventGet Get live events
func LiveEventGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/liveevent", gin.H{
		"systemTime":     str.NowString(),
		"liveeventviews": service.GetLiveEvent(),
	})
}
