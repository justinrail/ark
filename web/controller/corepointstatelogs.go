package controller

import (
	"ark/hub/domain"
	"ark/util/str"
	"ark/web/vm"
	"bytes"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//GetCorePointStatePacketByCorePointID 根据测点获取相应的事件包历史
func GetCorePointStatePacketByCorePointID(c *gin.Context) {
	items := make([]vm.LiteItem, 0)
	gID := c.Query("gatewayId")
	cID := c.Query("coresourceId")
	pID := c.Query("corepointId")
	pID = strings.TrimSuffix(pID, "/")

	gatewayID, err := strconv.Atoi(gID)
	coresourceID, err2 := strconv.Atoi(cID)
	corepointID, err3 := strconv.Atoi(pID)

	if err == nil && err2 == nil && err3 == nil {
		gateway, ok := domain.Gateways.Get(gatewayID)

		if ok {

			coresource, exist := (gateway.(*domain.Gateway)).CoreSources.Get(coresourceID)

			if exist {
				corepoint, here := (coresource.(*domain.CoreSource)).CorePoints.Get(corepointID)

				if here {
					plogs := (corepoint.(*domain.CorePoint)).StateLogs

					for e := plogs.Front(); e != nil; e = e.Next() {
						cov := (e.Value.(*domain.COV))

						item := vm.LiteItem{}
						buffer := bytes.NewBufferString("")
						buffer.WriteString("[E:")
						buffer.WriteString(strconv.Itoa(cov.StateID))
						buffer.WriteString("]")
						item.ItemName = buffer.String()
						item.ItemValue = str.TimeToString(cov.Timestamp)

						items = append(items, item)
					}
				}

			}
		}

	}

	c.JSON(200, items)
}
