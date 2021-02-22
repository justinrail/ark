package controller

import (
	"ark/hub/domain"
	"ark/web/vm"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//GetLitePoints 获取实时数据
func GetLitePoints(c *gin.Context) {
	items := make([]vm.LitePoint, 0)
	pidsParam := c.Query("corePointIds")
	cidsParam := c.Query("coreSourceIds")

	if len(pidsParam) == 0 && len(cidsParam) == 0 {
		c.JSON(200, items)
		return
	}

	if len(pidsParam) > 0 {
		ids := strings.Split(pidsParam, ",")

		if len(ids) == 0 {
			c.JSON(200, items)
			return
		}

		for _, id := range ids {
			idr, err := strconv.Atoi(id)
			if err == nil {
				cp, ok := domain.CorePoints.Get(idr)
				if ok {
					corePoint := cp.(*domain.CorePoint)
					item := vm.LitePoint{}

					item.ID = idr
					val, err := corePoint.GetCurrentValue()

					if err == nil {
						item.Val = val
					}

					items = append(items, item)
				}
			}
		}

		c.JSON(200, items)
	}

	if len(cidsParam) > 0 {
		cids := strings.Split(cidsParam, ",")

		if len(cids) == 0 {
			c.JSON(200, items)
			return
		}

		for _, cid := range cids {
			idr, err := strconv.Atoi(cid)
			if err == nil {
				cs, ok := domain.CoreSources.Get(idr)
				if ok {
					coreSource := cs.(*domain.CoreSource)

					cps := coreSource.CorePoints.Iter()

					for cp := range cps {
						corePoint := cp.Value.(*domain.CorePoint)
						item := vm.LitePoint{}

						item.ID = corePoint.CorePointID
						val, err := corePoint.GetCurrentValue()

						if err == nil {
							item.Val = val
						}

						items = append(items, item)
					}
				}
			}
		}

		c.JSON(200, items)
	}

}
