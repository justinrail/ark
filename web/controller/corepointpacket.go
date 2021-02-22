package controller

import (
	"ark/hub/domain"
	"ark/util/str"
	"ark/web/vm"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//GetCorePointPacketByCorePointID 根据测点获取相应的包历史
func GetCorePointPacketByCorePointID(c *gin.Context) {
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
					plogs := (corepoint.(*domain.CorePoint)).PacketLogs

					for e := plogs.Front(); e != nil; e = e.Next() {
						cov := (e.Value.(*domain.COV))

						item := vm.LiteItem{}
						buffer := bytes.NewBufferString("")
						buffer.WriteString("[N:")
						buffer.WriteString(fmt.Sprintf("%.2f", cov.CurrentNumericValue))
						buffer.WriteString("] \n[S:")
						buffer.WriteString(cov.CurrentStringValue)
						buffer.WriteString("] [V:")
						buffer.WriteString(strconv.FormatBool(cov.IsValid))
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
