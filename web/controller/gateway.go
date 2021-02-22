package controller

import (
	"ark/hub/domain"
	"ark/web/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetGateways 获取所有的网关FSU
func GetGateways(c *gin.Context) {
	collectorParam := c.Query("collector")

	if len(collectorParam) > 0 {
		c.JSON(200, service.GetGatewaysByCollector(collectorParam))
	} else {
		c.JSON(200, service.GetAllGateway())
	}
}

//UpdateGateway UpdateGateway
func UpdateGateway(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("gatewayId"))

	gw := &domain.Gateway{ID: id}
	err2 := c.BindJSON(gw)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong format",
		})
	}

	if err == nil && err2 == nil {
		c.JSON(http.StatusOK, service.UpdateGatewayJoined(id, gw.Joined))
	}
}

//DeleteGatewayByID 根据id删除gateway
func DeleteGatewayByID(c *gin.Context) {
	gwParam := c.Param("gatewayId")

	if len(gwParam) == 0 {
		c.JSON(404, gin.H{})
		return
	}
	idr, err := strconv.Atoi(gwParam)

	if err == nil {
		service.DeleteGatewayByGatewayID(idr)
		c.JSON(200, gin.H{})
		return
	}

	c.JSON(505, gin.H{})
}

//GetGatewayByID 获取的网关FSU
func GetGatewayByID(c *gin.Context) {
	items := make([]domain.Gateway, 0)
	gwParam := c.Param("gatewayId")

	if len(gwParam) == 0 {
		c.JSON(404, items)
		return
	}
	if len(gwParam) > 0 {

		idr, err := strconv.Atoi(gwParam)
		if err == nil {
			c.JSON(200, service.GetGatewayByID(idr))
		}
	}
}
