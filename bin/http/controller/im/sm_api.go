/**
  @author:panliang
  @data:2021/8/10
  @note
**/
package im

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_im/pkg/config"
	"go_im/pkg/redis"
	"go_im/pkg/response"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type SmApiController struct {}

var username = config.GetString("app.sm_name")
var password = config.GetString("app.sm_password")

type ResponseData struct {
	Success bool 	`json:"success"`
	Code string `json:"code"`
	Message string `json:"message"`
	Data Data `json:"data"`
	RequestId string `json:"RequestId"`
}
type Data struct {
	Token string `json:"token"`
}

func (*SmApiController) GetApiToken(c *gin.Context){

	stringCmd := redis.RedisDB.Get("sm_token")
	if len(stringCmd.Val()) !=0 {
		fmt.Println("字符串不为空")
		resp := new(ResponseData)
		resp.Code="success"

		resp.Data.Token=stringCmd.Val()
		resp.Success=true
		c.JSON(200,resp)
		return


	}


	data := url.Values{"username":{username},"password":{password}}
	
	j,err :=http.PostForm("https://sm.ms/api/v2/token",data)
	if err !=nil {
		fmt.Println(err)
	}
	defer j.Body.Close()

	bodyC, _ := ioutil.ReadAll(j.Body)
	resp := new(ResponseData)
	json.Unmarshal(bodyC,resp)
	if resp.Success {
		response.FailResponse(500,resp.Message)
		return
	}
	redis.RedisDB.Set("sm_token",resp.Data.Token,time.Hour*1)
	c.JSON(200,resp)
	
}