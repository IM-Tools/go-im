/**
  @author:panliang
  @data:2021/6/25
  @note
**/

package middleware

import (
	"github.com/gin-gonic/gin"
)

//跨域访问：cross  origin resource share
func CrosHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "false")
		context.Set("content-type", "application/json")

		if method == "OPTIONS" {
			//response.FailResponse(401, "Options Request!").ToJson(context)
		}
		//处理请求
		context.Next()
	}
	// return func(c *gin.Context) {
	// 	method := c.Request.Method

	// 	c.Header("Access-Control-Allow-Origin", "*")
	// 	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	// 	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	// 	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	// 	c.Header("Access-Control-Allow-Credentials", "true")

	// 	//放行所有OPTIONS方法
	// 	if method == "OPTIONS" {
	// 		c.AbortWithStatus(403)
	// 	}
	// 	// 处理请求
	// 	c.Next()
	// }
}
