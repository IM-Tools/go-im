/**
  @author:panliang
  @data:2021/9/16
  @note
**/
package cache

import (
	"encoding/json"
	"im_app/pkg/redis"
	"im_app/pkg/zaplog"
	"strconv"
)

func getGroupIdsStr(group_id int) string {
	return "core:group:" + strconv.Itoa(group_id)
}

// todo
func getGroup(group_id int) map[int]int {
	groupId := make(map[int]int)
	str := getGroupIdsStr(group_id)
	data := redis.RedisDB.Get(str)
	if len(data.Val()) > 0 {
		by_data, err := data.Bytes()
		if err != nil {
			zaplog.Error("----获取群组用户id失败", err)
		}
		json.Unmarshal(by_data, groupId)
	} else {

	}

	return groupId
}
