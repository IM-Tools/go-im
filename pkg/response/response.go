/**
  @author:panliang
  @data:2021/6/24
  @note
**/
package response

import "github.com/gin-gonic/gin"

const (
	SuccessCode = 200
	ErrorCode = 500
)

//响应结构体
type Response struct {
	Code  int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//响应json
func (resp *Response) ToJson(ctx *gin.Context) {
	code := 200
	if resp.Code != SuccessCode {
		code = resp.Code
	}
	ctx.JSON(code, resp)
}

//失败响应
func FailResponse(code int,message string) *Response    {
	return &Response{
		Code:  ErrorCode,
		Message: message,
		Data:    nil,
	}
}
//成功响应
func SuccessResponse(data interface{},code int) *Response  {
	return &Response{
		Code:  code,
		Message: "Success",
		Data:    data,
	}
}
