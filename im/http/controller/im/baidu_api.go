/**
  @author:panliang
  @data:2021/8/12
  @note
**/
package im

import (
	"github.com/gin-gonic/gin"
	"go_im/im/utils"
	"go_im/pkg/config"
	"go_im/pkg/response"
)

type BaiduController struct {}
var ym = config.GetString("app.ym")
func (*BaiduController) UploadVoiceFile(c *gin.Context)  {
	voice, _ := c.FormFile("voice")
	dir := utils.GetCurrentDirectory()
	// 上传文件至指定目录 没找到第三方免费的第三方存储 先用自己的吧
	path :=dir+"/voice/"+voice.Filename
	 c.SaveUploadedFile(voice, path)
	response.SuccessResponse(map[string]interface{}{
		"url":ym+"voice/"+voice.Filename,
	},200).ToJson(c)
	return
}

