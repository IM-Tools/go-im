/**
  @author:panliang
  @data:2021/8/12
  @note
**/
package im

import (
	"github.com/gin-gonic/gin"
	"go_im/pkg/response"
)

type BaiduController struct {

}
func (*BaiduController) UploadVoiceFile(c *gin.Context)  {
	voice, _ := c.FormFile("voice")
	//dir := utils.GetCurrentDirectory()
	// 上传文件至指定目录 没找到第三方免费的第三方存储 先用自己的吧
	path :="/www/wwwroot/qc_admin.pltrue.top/public/voice/"+voice.Filename
	 c.SaveUploadedFile(voice, path)
	response.SuccessResponse(map[string]interface{}{
		"url":"http://voice.pltrue.top/voice/"+voice.Filename,
	},200).ToJson(c)
	return
}

