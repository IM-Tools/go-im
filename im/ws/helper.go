/**
  @author:panliang
  @data:2021/7/2
  @note
**/
package ws

import (
	"encoding/json"
	"fmt"
	"go_im/pkg/model"
)

//byte -> map
func EnMessage(message []byte) (msg *Message) {
	err := json.Unmarshal([]byte(string(message)),&msg)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return
}

//get chat group user id
func GetGroupUid(group_id int) ([]GroupId,error) {
	var groups []GroupId
	err := model.DB.Table("im_group_users").Where("group_id=?",group_id).Find(&groups).Error;if err != nil {
		return groups,err
	}
	return groups,nil
}



