package controller

import (
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/str"
	"ark/web/vm"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetCoreSourcePacketByCoreSourceID 根据网关获取相应的包历史
func GetCoreSourcePacketByCoreSourceID(c *gin.Context) {
	items := make([]vm.LiteItem, 0)
	gID := c.Query("gatewayId")
	cID := c.Query("coresourceId")

	gatewayID, err := strconv.Atoi(gID)
	coresourceID, err2 := strconv.Atoi(cID)

	if err == nil && err2 == nil {
		gateway, ok := domain.Gateways.Get(gatewayID)

		if ok {

			coresource, exist := (gateway.(*domain.Gateway)).CoreSources.Get(coresourceID)

			if exist {
				plogs := (coresource.(*domain.CoreSource)).PacketLogs

				for e := plogs.Front(); e != nil; e = e.Next() {
					cor := (e.Value.(*domain.COR))

					item := vm.LiteItem{}
					item.ItemName = enum.GetEnumString(cor.Flag)
					item.ItemValue = str.TimeToString(cor.Timestamp)

					items = append(items, item)
				}

			}
		}

	}

	c.JSON(200, items)
}
