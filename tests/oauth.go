/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package main

import (
	"encoding/json"
	"fmt"
	"go_im/pkg/helpler"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main()  {

	urls :="http://adminapi.test/api/captcha"
	data := url.Values{"app_id":{"238b2213-a8ca-42d8-8eab-1f1db3c50ed6"}, "mobile_tel":{"13794227450"}}
	body := strings.NewReader(data.Encode())
	resp,err := http.Post(urls,"application/x-www-form-urlencoded",body)

	if err!=nil{
		fmt.Println(err)
	}

	defer resp.Body.Close()

	bodyC, _ := ioutil.ReadAll(resp.Body)

	jsonMap := helpler.JsonToMap(bodyC)

	err = json.Unmarshal(bodyC, &jsonMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	datas := jsonMap["data"]
	fmt.Println(datas)

}
