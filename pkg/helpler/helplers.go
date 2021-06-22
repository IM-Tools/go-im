/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package helpler

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/url"
)

// json string to a map type
func JsonToMap(str []byte) map[string]interface{}  {
	var jsonMap  map[string]interface{}
	err := json.Unmarshal(str, &jsonMap)
	if err != nil {
		panic(err)
	}
	return jsonMap
}

func HttpBuildQuery(queryData url.Values) string {
	return queryData.Encode()
}

//加密算法
func HashAndSalt(pwd string) string {
	hash,err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//解密算法
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	plainPwds := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwds)
	if err != nil {
		return false
	}
	return true
}


