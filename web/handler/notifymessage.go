package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//NotifyMessageGet Get notify messages
func NotifyMessageGet(c *gin.Context) {

	ruleName := c.Query("notifyruleName")
	queueID := c.Query("messagequeueId")

	c.HTML(http.StatusOK, "view/notifymessage", gin.H{
		"systemTime":     str.NowString(),
		"NotifyRuleName": ruleName,
		"messages":       service.GetNotifyMessage(queueID),
	})
}
