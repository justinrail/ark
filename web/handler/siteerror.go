package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Error404Get HTTP 404 Page
func Error404Get(c *gin.Context) {

	//根目录文件要以tpl结尾并全名书写
	c.HTML(http.StatusOK, "404.tpl", nil)

}
