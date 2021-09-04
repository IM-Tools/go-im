/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package utils

import (
	"github.com/tidwall/gjson"
	"go_im/im/helper"
	"go_im/pkg/config"
	"go_im/pkg/helpler"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var client_id = config.GetString("oauth.wb_client_id")
var client_secret = config.GetString("oauth.wb_client_secret")
var redirect_uri = config.GetString("oauth.wb_redirect_uri")
var access_token_url = "https://api.weibo.com/oauth2/access_token"
var user_info_url = "https://api.weibo.com/2/users/show.json"
var get_token_info="https://api.weibo.com/oauth2/get_token_info"


// Result represents a json value that is returned from GetUserInfo().

type UserInfo struct {
	Name string
	Email string
	Avatar string
	OauthId string
	BoundOauth int
}

// GetAccessToken function string returns an string access_token.str

func GetWeiBoAccessToken(code *string) string  {
	queryData :=url.Values{"client_id":{client_id},
		"code":{*code},
		"client_secret":{client_secret},
		"redirect_uri":{redirect_uri},
		"grant_type":{"authorization_code"}}

	urls :=access_token_url+"?"+helpler.HttpBuildQuery(queryData)

	data := url.Values{}
	body := strings.NewReader(data.Encode())
	resp,err := http.Post(urls,"application/x-www-form-urlencoded",body)
	if err!=nil{
		helper.Error("获取code失败",err)
	}
	defer resp.Body.Close()

	bodyC, _ := ioutil.ReadAll(resp.Body)

	access_token := gjson.Get(string(bodyC),"access_token")

	return access_token.Str
}

// GetUserInfo function  returns an UserInfo

func GetWeiBoUserInfo(access_token *string) string {

	uid :=getUid(&*access_token)

	urls := user_info_url+"?uid="+ uid+"&access_token="+*access_token
	resp,err := http.Get(urls)
	if err!=nil{
		helper.Error("请求微博获取access_token失败",err)
	}
	defer resp.Body.Close()

	bodyC, _ := ioutil.ReadAll(resp.Body)

	return string(bodyC)

}
// get uid
func getUid(access_token *string) string  {
	urls := get_token_info+"?access_token="+*access_token
	data := url.Values{}
	body := strings.NewReader(data.Encode())
	resp,err := http.Post(urls,"application/x-www-form-urlencoded",body)
	if err!=nil{
		helper.Error("请求微博获取uid失败",err)
	}
	defer resp.Body.Close()

	bodyC, _ := ioutil.ReadAll(resp.Body)

	uid := gjson.Get(string(bodyC),"uid")

	return uid.Raw
}
