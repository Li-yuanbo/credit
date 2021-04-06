package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回请求信息
func SucResponse(c *gin.Context, data interface{}){
	c.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  "SUCCESS",
		"data":	data,
	})
}

//返回错误信息
func ErrResponse(c *gin.Context, code int, errCode int, msg string){
	c.JSON(code, gin.H{
		"code": errCode,
		"msg":	msg,
	})
}