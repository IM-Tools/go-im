/**
  @author:panliang
  @data:2021/8/16
  @note
**/
package zaplog

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

//全局异常处理
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			errors := errorToString(r)
			//写入日志
			Warning(errors)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器异常",
				"data": nil,
			})
			c.Abort()
		}
	}()
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
