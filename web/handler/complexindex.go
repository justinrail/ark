package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ComplexIndexGet Get complexIndex
func ComplexIndexGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/complexindex", gin.H{
		"systemTime":    str.NowString(),
		"complexindexs": service.GetAllComplexIndexView(),
	})
}
