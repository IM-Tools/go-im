/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package helpler

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/url"
	"time"
)

// json string to a map type
func JsonToMap(str []byte) map[string]interface{} {
	var jsonMap map[string]interface{}
	err := json.Unmarshal(str, &jsonMap)
	if err != nil {
		panic(err)
	}
	return jsonMap
}

func HttpBuildQuery(queryData url.Values) string {
	return queryData.Encode()
}

func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	plainPwds := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwds)
	if err != nil {
		return false
	}
	return true
}

func ProduceChannelName(f_id int64, t_id int64) (channel_a string, channel_b string) {
	channel_a = fmt.Sprintf("channel_%v_%v", f_id, t_id)
	channel_b = fmt.Sprintf("channel_%v_%v", t_id, f_id)
	return channel_a, channel_b
}
func ProduceChannelGroupName(t_id string) string {
	return "channel_" + t_id
}

func GetNowFormatTodayTime() string {

	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()),
		now.Day())

	return dateStr
}

func CreateEmailCode() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}
