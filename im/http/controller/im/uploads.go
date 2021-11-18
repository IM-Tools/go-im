/**
  @author:panliang
  @data:2021/8/12
  @note
**/
package im

import (
	"github.com/gin-gonic/gin"
	"im_app/im/utils"
	"im_app/pkg/config"
	"im_app/pkg/response"
)

type UploadController struct{}

var ym = config.GetString("app.ym")

// @BasePath /api

// @Summary 音频文件上传接口
// @Description 音频文件上传接口
// @Tags 音频文件上传接口
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param voice formData file true "图片上传"
// @Produce json
// @Success 200
// @Router /UploadVoiceFile [post]
func (*UploadController) UploadVoiceFile(c *gin.Context) {
	voice, _ := c.FormFile("voice")
	dir := utils.GetCurrentDirectory()
	// 上传文件至指定目录 没找到第三方免费的第三方存储 先用自己的吧
	path := dir + "/voice/" + voice.Filename
	c.SaveUploadedFile(voice, path)
	response.SuccessResponse(map[string]interface{}{
		"url": ym + "voice/" + voice.Filename,
	}, 200).ToJson(c)
	return
}
