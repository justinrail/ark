package handler

import (
	"ark/util/str"
	"net/http"

	"github.com/gin-gonic/gin"
)

//IndexGet Web Path / handler
func IndexGet(c *gin.Context) {

	c.HTML(http.StatusOK, "view/home", gin.H{
		"systemTime": str.NowString(),
	})

}
