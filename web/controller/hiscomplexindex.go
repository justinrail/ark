package controller

import (
	"ark/hub/domain"
	"ark/util/str"
	"ark/web/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//GetHisComplexIndexs 获取历史指标数据
func GetHisComplexIndexs(c *gin.Context) {
	items := make([]domain.HisComplexIndex, 0)
	cids := make([]int, 0)
	cidsParam := c.Query("complexIndexIds")
	startParam := c.Query("startTime")
	endParam := c.Query("endTime")

	if len(cidsParam) == 0 {
		c.JSON(404, items)
		return
	}

	if len(cidsParam) > 0 {
		ids := strings.Split(cidsParam, ",")

		if len(ids) == 0 {
			c.JSON(404, items)
			return
		}

		for _, id := range ids {
			idr, err := strconv.Atoi(id)
			if err == nil {
				cids = append(cids, idr)
			}
		}

		if len(startParam) > 0 && len(endParam) > 0 {
			startTime := str.LocalTimstampStringToTime(startParam)
			endTime := str.LocalTimstampStringToTime(endParam)
			c.JSON(200, service.GetHisComplexIndexByIDsANDTimeRange(cids, startTime, endTime))
		} else {
			c.JSON(200, service.GetHisComplexIndexByIDs(cids))
		}
	}

}
