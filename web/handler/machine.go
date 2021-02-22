package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//MachineGet Get machine data to page
func MachineGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/machine", gin.H{
		"systemTime": str.NowString(),
		"machines":   service.GetAllMachine(),
	})

}
