/**
  @author:panliang
  @data:2021/10/29
  @note
**/
package dao

import (
	"go_im/im/http/models/friend"
	"go_im/pkg/model"
	"time"
)

type UserService struct{}

//添加默认好友关系
func (*UserService) AddDefaultFriend(m_id uint64) {
	model.DB.Create(&friend.ImFriends{FId: m_id, MId: 1, Status: 1, CreatedAt: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")})

	model.DB.Create(&friend.ImFriends{FId: 1, MId: m_id, Status: 1, CreatedAt: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")})

}
