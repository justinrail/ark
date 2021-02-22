package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//NotifyRuleGet Get NotifyRuleGet
func NotifyRuleGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/notifyrule", gin.H{
		"systemTime":  str.NowString(),
		"notifyrules": service.GetAllNotifyRule(),
	})
}
