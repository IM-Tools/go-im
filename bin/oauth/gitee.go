/**
  @author:panliang
  @data:2021/6/25
  @note
**/
package oauth

import (
	"fmt"
	"go_im/pkg/config"
	"io/ioutil"
	"net/http"
	"net/url"
)

var ge_client_id = config.GetString("oauth.ge_client_id")
var ge_client_secret = config.GetString("oauth.ge_client_secret")
var ge_redirect_uri = config.GetString("oauth.ge_redirect_uri")

const (
	get_token_url ="https://gitee.com/oauth/token"
	get_user_url ="https://gitee.com/api/v5/use"
)

func GetGiteeAccessToken(code *string)  string {

	fmt.Println(code)
	queryData :=url.Values{
		"client_id":     {ge_client_id},
		"redirect_uri":  {ge_redirect_uri},
		"code":          {*code},
		"grant_type":    {"authorization_code"},
		"client_secret": {ge_client_secret},
	}
	resp,err := http.PostForm(get_token_url,queryData)
	if err!=nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	bodyC, _ := ioutil.ReadAll(resp.Body)

	return string(bodyC)
//	access_token := gjson.Get(string(bodyC),"access_token")
}

func GetGiteeUserInfo(access_token *string) string  {
	urls := get_user_url+"?access_token="+*access_token
	resp,err := http.Get(urls)
	if err!=nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	bodyC, _ := ioutil.ReadAll(resp.Body)

	return string(bodyC)
}