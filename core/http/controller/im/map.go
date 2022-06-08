/**
  @author:panliang
  @data:2021/12/11
  @note
**/
package im

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"im_app/core/http/services"
	"im_app/pkg/helpler"
	"im_app/pkg/response"
)

type MapController struct {
}

//获取经纬度位置信息

func (*MapController) GetLongitude(cxt *gin.Context) {
	ip := helpler.GetLocalIP()
	fmt.Println(ip)
	service := new(services.MapService)
	result := service.GetLongitude(ip)
	response.SuccessResponse(result).ToJson(cxt)
	return
}
