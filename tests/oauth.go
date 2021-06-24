/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package main

import (
	"encoding/json"
	"fmt"
	//"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main()  {


	keys := httpReuqet();
	fmt.Println(keys)

}

type Response struct {
	// Code    int    `json:"code"`
	// Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Captcha Captcha `json:"captcha"`
}

type Captcha struct {
	Key       string `json:"key"`
	// Img       string `json:"img"`
	// Sensitive bool   `json:"senstive"`
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

	response := new(Response)

	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	bodyC, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyC, response)

	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return response.Data.Captcha.Key
	//bodyS := string(bodyC);
	//
	//keys := gjson.Get(bodyS,"data.captcha.key")

	//return keys.Str
}
