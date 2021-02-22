package handler

import (
	"ark/util/str"
	"ark/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//DomainGet Get domain root gateways
func DomainGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/domain", gin.H{
		"systemTime": str.NowString(),
		"gateways":   service.GetAllGateway(),
	})

}
