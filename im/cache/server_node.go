/**
  @author:panliang
  @data:2021/11/13
  @note
**/
package cache

import (
	"strconv"

	"im_app/pkg/config"
	"im_app/pkg/redis"
)

var (
	cache_node = config.GetString("app.node") + ":" + config.GetString("app.grpc_port")
)

// 必须有一个结构体去实现该接口的方法

type ServiceNodeHandler interface {
	SetUserServiceNode(ID int64)        // 设置用户节点
	GetUserServiceNode(ID int64) string // 获取用户节点
	DelUserServiceNode(ID int64)        // 删除用户节点
}

type ServiceNode struct {
}

func getUserIdStr(ID int64) string {
	// 注意 ：写法 在可视化工具里面可以更好的看到缓存的结构体
	return "im:node:user:" + strconv.Itoa(int(ID))
}

func (node *ServiceNode) GetUserServiceNode(ID int64) string {
	var key = getUserIdStr(ID)
	StringCmd := redis.RedisDB.Get(key)
	return StringCmd.Val()
}

func (node *ServiceNode) SetUserServiceNode(ID int64) {
	var key = getUserIdStr(ID)
	var value = cache_node
	redis.RedisDB.Set(key, value, 0)
}

func (node *ServiceNode) DelUserServiceNode(ID int64) {
	var key = getUserIdStr(ID)
	redis.RedisDB.Del(key)
}
