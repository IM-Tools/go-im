/**
  @author:panliang
  @data:2021/11/13
  @note
**/
package cache

import (
	"im_app/pkg/config"
	"im_app/pkg/redis"
	"strconv"
)

var (
	cache_node = config.GetString("app.node") + ":" + config.GetString("app.grpc_port")
)

type ServiceNodeHandler interface {
	SetUserServiceNode(ID int64)
	GetUserServiceNode(ID int64) string
	DelUserServiceNode(ID int64)
}

func getUserIdStr(ID int64) string {
	return "im:node:user:" + strconv.Itoa(int(ID))
}

func GetUserServiceNode(ID int64) string {
	var key = getUserIdStr(ID)
	StringCmd := redis.RedisDB.Get(key)
	return StringCmd.Val()
}

func SetUserServiceNode(ID int64) {
	var key = getUserIdStr(ID)
	var value = cache_node
	redis.RedisDB.Set(key, value, 0)
}

func DelUserServiceNode(ID int64) {
	var key = getUserIdStr(ID)
	redis.RedisDB.Del(key)
}
