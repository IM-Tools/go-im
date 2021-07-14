/**
  @author:panliang
  @data:2021/7/9
  @note
**/
package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var access_token = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89"

func Upload(base64 string,path string,message string)  string {

	urls :="https://gitee.com/api/v5/repos/pltrue/figure-bed/contents/"+path

	resp,err := http.PostForm(urls, map[string][]string{
		"access_token":{access_token},
		"content":{base64},
		"message":{message},
	})
	if err !=nil{
		fmt.Println(err)
	}

	defer resp.Body.Close()


	bodyC, _ := ioutil.ReadAll(resp.Body)

	return string(bodyC)

}
