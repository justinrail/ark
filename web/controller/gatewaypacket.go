package controller

import (
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/str"
	"ark/web/vm"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetGatewayPacketByGatewayID 根据网关获取相应的包历史
func GetGatewayPacketByGatewayID(c *gin.Context) {
	items := make([]vm.LiteItem, 0)
	id := c.Param("gatewayId")

	idr, err := strconv.Atoi(id)

	if err == nil {
		gateway, ok := domain.Gateways.Get(idr)

		if ok {
			plogs := (gateway.(*domain.Gateway)).PacketLogs

			for e := plogs.Front(); e != nil; e = e.Next() {
				cog := (e.Value.(*domain.COG))

				item := vm.LiteItem{}
				item.ItemName = enum.GetEnumString(cog.Flag)
				item.ItemValue = str.TimeToString(cog.Timestamp)

				items = append(items, item)
			}
		}

	}
	// if ok {
	// 	css := (gateway.(*domain.Gateway)).CoreSources.Iter()

	// 	for cs := range css {
	// 		c := (cs.Value.(*domain.CoreSource))

	// 		item := vm.LiteItem{}
	// 		item.ItemID = c.

	// 		items = append(items, item)

	// 	}
	// }

	//fmt.Println(id)

	c.JSON(200, items)
}
