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
	"go_im/im/utils"
	"go_im/pkg/config"
	"go_im/pkg/redis"
	"go_im/pkg/response"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type SmApiController struct{}

var username = config.GetString("app.sm_name")
var password = config.GetString("app.sm_password")
var sm_token = config.GetString("app.sm_token")

type ResponseData struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	Data      Data   `json:"data"`
	RequestId string `json:"RequestId"`
}
type Data struct {
	Token string `json:"token"`
}

func (*SmApiController) GetApiToken(c *gin.Context) {
	stringCmd := redis.RedisDB.Get("sm_token")
	if len(stringCmd.Val()) != 0 {

		resp := new(ResponseData)
		resp.Code = "success"
		resp.Data.Token = stringCmd.Val()
		resp.Success = true
		fmt.Println(resp)
		c.JSON(200, resp)
		return
	}
	data := url.Values{"username": {username}, "password": {password}}
	j, err := http.PostForm("https://sm.ms/api/v2/token", data)
	utils.LogError(err)
	defer j.Body.Close()
	bodyC, _ := ioutil.ReadAll(j.Body)
	resp := new(ResponseData)
	json.Unmarshal(bodyC, resp)
	if resp.Success {
		response.FailResponse(500, resp.Message)
		return
	}
	redis.RedisDB.Set("sm_token", resp.Data.Token, time.Hour*1)
	c.JSON(200, resp)
}

type ResponseUploadData struct {
	Success   bool        `json:"'success'"`
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Data      DataSuccess `json:"data"`
	RequestId string      `json:"RequestId"`
}

type DataSuccess struct {
	FileId    int    `json:"file_id"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Filename  string `json:"filename"`
	Storename string `json:"storename"`
	Size      int    `json:"size"`
	Path      string `json:"path"`
	Hash      string `json:"hash"`
	Url       string `json:"url"`
	Delete    string `json:"delete"`
	Page      string `json:"page"`
}

func (*SmApiController) UploadImg(c *gin.Context) {
	file, _ := c.FormFile("Smfile")
	dir := utils.GetCurrentDirectory()
	path :=dir+"/docs/"+file.Filename
	err := c.SaveUploadedFile(file, path)
	utils.LogError(err)
	header := new(utils.Header)
	header.Authorization = "Authorization"
	header.Token = sm_token
	resp, err := utils.PostFile(path, "https://sm.ms/api/v2/upload", header)
	utils.LogError(err)
	bodyC, _ := ioutil.ReadAll(resp.Body)
	data := new(ResponseUploadData)
	json.Unmarshal(bodyC, data)
	c.JSON(200, data)
}

