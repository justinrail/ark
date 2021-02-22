package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GatewayGet Get Coresources of gateway x
func GatewayGet(c *gin.Context) {

	id := c.Param("gatewayId")

	idr, err := strconv.Atoi(id)

	if err == nil {
		c.HTML(http.StatusOK, "view/gateway", gin.H{
			"systemTime":  str.NowString(),
			"coresources": service.GetCoreSourceByGateway(idr),
		})
	}
}

//DeleteGatewayByGatewayID delete gateway
func DeleteGatewayByGatewayID(c *gin.Context) {

	id := c.Query("gatewayId")

	idr, err := strconv.Atoi(id)

	if err == nil {
		service.DeleteGatewayByGatewayID(idr)

		c.HTML(http.StatusOK, "view/domain", gin.H{
			"systemTime": str.NowString(),
			"gateways":   service.GetAllGateway(),
		})
	}
}
