package controller

import (
	"ark/web/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetComplexIndexs 根据获取指标及当前值
func GetComplexIndexs(c *gin.Context) {
	id := c.Param("complexIndexId")

	if len(id) == 0 {
		c.JSON(200, service.GetAllComplexIndex())
	} else {
		idr, err := strconv.Atoi(id)
		if err == nil {
			c.JSON(200, service.GetComplexIndexByID(idr))
		} else {
			c.JSON(404, nil)
		}
	}
}
