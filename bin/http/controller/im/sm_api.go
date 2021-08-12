/**
  @author:panliang
  @data:2021/8/10
  @note
**/
package im

import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_im/bin/utils"
	"go_im/pkg/config"
	"go_im/pkg/redis"
	"go_im/pkg/response"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type SmApiController struct {}

var username = config.GetString("app.sm_name")
var password = config.GetString("app.sm_password")
var sm_token = config.GetString("app.sm_token")

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

		resp := new(ResponseData)
		resp.Code="success"
		resp.Data.Token=stringCmd.Val()
		resp.Success=true
		fmt.Println(resp)
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

type ResponseUploadData struct {
	Success bool `json:"'success'"`
	Code string `json:"code"`
	Message string `json:"message"`
	Data DataSuccess `json:"data"`
	RequestId string `json:"RequestId"`
}

type DataSuccess struct {
	FileId int `json:"file_id"`
	Width int `json:"width"`
	Height int `json:"height"`
	Filename string `json:"filename"`
	Storename string `json:"storename"`
	Size int `json:"size"`
	Path string `json:"path"`
	Hash string `json:"hash"`
	Url string `json:"url"`
	Delete string `json:"delete"`
	Page string `json:"page"`
}

func (*SmApiController) UploadImg(c *gin.Context)  {
	file, _ := c.FormFile("smfile")

	dir := getCurrentDirectory()
	fmt.Println(dir)
	path :=dir+"/docs/"+file.Filename
	err := c.SaveUploadedFile(file,path)
	if err != nil {
		fmt.Println(err)
	}
	header := new(utils.Header)
	header.Authorization = "Authorization"
	header.Token = sm_token
	resp,err := utils.PostFile(path,"https://sm.ms/api/v2/upload",header)
	if err != nil {
		fmt.Println(err)
	}
	bodyC, _ := ioutil.ReadAll(resp.Body)
	data := new(ResponseUploadData)
	json.Unmarshal(bodyC,data)
    c.JSON(200,data)
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir //strings.Replace(dir, "\\", "/", -1)
}

