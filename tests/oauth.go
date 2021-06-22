/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"go_im/pkg/helpler"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main()  {


	//keys := httpReuqet();
	//fmt.Println(keys)

}


func httpReuqet() string  {

	urls :="http://adminapi.test/api/captcha"
	data := url.Values{"app_id":{""}, "mobile_tel":{""}}
	body := strings.NewReader(data.Encode())
	resp,err := http.Post(urls,"application/x-www-form-urlencoded",body)

	if err!=nil{
		fmt.Println(err)
	}

	defer resp.Body.Close()

	bodyC, _ := ioutil.ReadAll(resp.Body)


	jsonMap := helpler.JsonToMap(bodyC)

	err = json.Unmarshal(bodyC, &jsonMap)

	bodyS := string(bodyC);

	keys := gjson.Get(bodyS,"data.captcha.key")

	return keys.Str
}
