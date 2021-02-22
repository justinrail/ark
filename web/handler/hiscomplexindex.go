package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//HisComplexIndexGet Get complexIndex
func HisComplexIndexGet(c *gin.Context) {
	//id := c.Param("complexindexId")
	id := c.Query("complexindexId")
	idr, err := strconv.Atoi(id)

	if err == nil {
		c.HTML(http.StatusOK, "view/hiscomplexindex", gin.H{
			"systemTime":       str.NowString(),
			"ComplexIndexID":   id,
			"hiscomplexindexs": service.GetHisComplexIndexViewByID(idr),
		})
	}
}
