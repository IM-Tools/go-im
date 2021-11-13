/**
  @author:panliang
  @data:2021/6/24
  @note
**/
package response

import "github.com/gin-gonic/gin"

/**
公共响应

统一定义如何将 struct 映射到响应
统一定义应用状态码
统一处理应用状态码与 http 状态码映射关系
*/
const (
	MaskNeedAuthor   = 8
	MaskParamMissing = 7
	StatusSuccess    = 200
	StatusError      = 200
)

//响应结构体
type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//响应json
func (resp *JsonResponse) ToJson(ctx *gin.Context) {
	code := 200
	if resp.Code != StatusSuccess {
		code = resp.Code
	}
	ctx.JSON(code, resp)
}

//失败响应
func FailResponse(code int, message string) *JsonResponse {
	return &JsonResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func SuccessResponse(data ...interface{}) *JsonResponse {
	var r interface{}
	if len(data) > 0 {
		r = data[0]
	}
	return &JsonResponse{
		Code:    StatusSuccess,
		Message: "Success",
		Data:    r,
	}
}

func ErrorResponse(status int, message string, data ...interface{}) *JsonResponse {
	return &JsonResponse{
		Code:    status,
		Message: message,
		Data:    data,
	}
}

// 将 json 设为响应体.
// HTTP 状态码由应用状态码决定
func (that *JsonResponse) WriteTo(ctx *gin.Context) {
	code := 200
	if that.Code != StatusSuccess {
		code = that.responseCode()
	}
	ctx.JSON(code, that)
}

// 获取 HTTP 状态码. HTTP 状态码由 应用状态码映射
func (that *JsonResponse) responseCode() int {
	// todo 完善应用状态码对应 http 状态码
	if that.Code != StatusSuccess {
		return 200
	}
	return 200
}
