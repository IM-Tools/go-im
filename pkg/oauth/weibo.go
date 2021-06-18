/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package oauth

import (
	"encoding/json"
	"fmt"
	"go_im/pkg/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


//获取access token

func getAccessToken(code string) (Arrays map[string]interface{}) {


	client_id     :=config.GetString("oauth.client_id")
	client_secret :=config.GetString("oauth.client_secret")
	redirect_uri  :=config.GetString("oauth.redirect_uri")

	access_token_urls := "https://api.weibo.com/oauth2/access_token"+"?client_id"+client_id+"&code="+code+"&client_secret"+client_secret+"&redirect_uri="+redirect_uri+"&grant_type=authorization_code"
	data := url.Values{"app_id":{"238b2213-a8ca-42d8-8eab-1f1db3c50ed6"}, "mobile_tel":{"13794227450"}}
	body := strings.NewReader(data.Encode())
	resp, err := http.Post(access_token_urls,"application/x-www-form-urlencoded",body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	bodyC, _ := ioutil.ReadAll(resp.Body)

	var jsonMap map[string]interface{}
	err = json.Unmarshal(bodyC, &jsonMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	return jsonMap
}

//func GetUserInfo(access_token string) (user map[string]interface{})  {
//
//}
//
//func getUid(access_token string)  {
//
//	url :="https://api.weibo.com/oauth2/get_token_info?access_token="+access_token
//
//	data := url.Values{"app_id":{"238b2213-a8ca-42d8-8eab-1f1db3c50ed6"}, "mobile_tel":{"13794227450"}}
//	body := strings.NewReader(data.Encode())
//	resp, err := http.Post(access_token_urls,"application/x-www-form-urlencoded",body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	defer resp.Body.Close()
//
//
//	bodyC, _ := ioutil.ReadAll(resp.Body)
//
//	var jsonMap map[string]interface{}
//	err = json.Unmarshal(bodyC, &jsonMap)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//}
