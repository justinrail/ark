package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//CoreSourceGet Get coresource by id
func CoreSourceGet(c *gin.Context) {

	gID := c.Query("gatewayId")
	cID := c.Query("coresourceId")
	gatewayID, err := strconv.Atoi(gID)
	coresourceID, err2 := strconv.Atoi(cID)

	if err == nil && err2 == nil {
		c.HTML(http.StatusOK, "view/coresource", gin.H{
			"systemTime":     str.NowString(),
			"GatewayID":      gID,
			"corepointviews": service.GetCorePointByCoreSource(gatewayID, coresourceID),
		})
	}

}
