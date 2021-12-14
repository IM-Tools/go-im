/**
  @author:panliang
  @data:2021/6/21
  @note
**/
package user

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"im_app/pkg/model"
	"time"
)

type Users struct {
	ID              int64  `json:"id"`
	Email           string `valid:"email" json:"email"`
	Password        string `valid:"password"`
	Avatar          string `json:"avatar"`
	Name            string `json:"name" valid:"name"`
	OauthType       int    `json:"oauth_type"`
	Status          int    `json:"status"`
	OauthId         string `json:"oauth_id"`
	CreatedAt       string `json:"created_at"`
	PasswordConfirm string ` gorm:"-" valid:"password_confirm"`
	LastLoginTime   string `json:"last_login_time"`
	Bio             string `json:"bio"`
	Sex             int    `json:"sex"`
	ClientType      int    `json:"client_type"`
	Age             int    `json:"age"`
}

type UserLists struct {
	ID            int64  `json:"id"`
	Email         string `valid:"email" json:"email"`
	Avatar        string `json:"avatar"`
	Name          string `json:"name" valid:"name"`
	Status        int    `json:"status"`
	CreatedAt     string `json:"created_at"`
	LastLoginTime string `json:"last_login_time"`
	Bio           string `json:"bio"`
	TopStatus     int    `json:"top_status"`
	Note          string `json:"note"`
}

type UsersWhiteList struct {
	ID            int64  `json:"id"`
	Email         string `valid:"email" json:"email"`
	Avatar        string `json:"avatar"`
	Name          string `json:"name" valid:"name"`
	OauthType     int    `json:"oauth_type"`
	Status        int    `json:"status"`
	OauthId       string `json:"oauth_id"`
	CreatedAt     string `json:"created_at"`
	Bio           string `json:"bio"`
	Sex           int    `json:"sex"`
	ClientType    int    `json:"client_type"`
	Age           int    `json:"age"`
	LastLoginTime string `json:"last_login_time"`
}

type ImFriendRecords struct {
	ID          int64  `json:"id"`
	FId         int64  `json:"f_id"`
	Status      int64  `json:"status"`
	CreatedAt   string `json:"created_at"`
	Information string `json:"information"`
}

// 字段过滤机制
func (u *Users) MarshalJSON() ([]byte, error) {

	// 将 User 的数据映射到 UsersWhiteList 上
	user := UsersWhiteList{
		ID:            u.ID,
		Email:         u.Email,
		Name:          u.Name,
		Avatar:        u.Avatar,
		CreatedAt:     u.CreatedAt,
		Bio:           u.Bio,
		Sex:           u.Sex,
		ClientType:    u.ClientType,
		Status:        u.Status,
		Age:           u.Age,
		LastLoginTime: u.LastLoginTime,
	}
	return json.Marshal(user)
}

func (Users) TableName() string {
	return "im_users"
}

// 当前登录用户
var AuthUser *Users

func (a *Users) AfterFind(tx *gorm.DB) (err error) {
	if a.Avatar == "" {
		a.Avatar = "https://cdn.learnku.com/uploads/avatars/27407_1531668878.png!/both/100x100"
	}
	return
}

func GetFriendListV2(user_id int64) ([]UserLists, error) {
	var users []UserLists

	sql := fmt.Sprintf("select u.id,u.email,u.avatar,u.name,u.status,"+
		"u.bio,u.client_type,u.last_login_time,f.status as top_status,f.note "+
		"from im_users as u left join im_friends as f on f.f_id=u.id"+
		" where f.m_id=%d order by f.status desc,f.top_time desc", user_id)

	err := model.DB.Raw(sql).Scan(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

// 设置用户上下线状态
func SetUserStatus(id int64, status int) {
	model.DB.Model(&Users{}).Where("id=?", id).Updates(Users{Status: status,
		LastLoginTime: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")})
}

func GetNotFriendList(subQuery *gorm.DB, id int64, name string) (userList []Users, err error) {

	var UserList []Users
	if len(name) > 0 {
		err = model.DB.
			Where("id not in (?) and id != ? and name like ?", subQuery, id, "%"+name+"%").
			Limit(10).
			Find(&userList).Error
	} else {
		err = model.DB.
			Where("id not in (?) and id != ?", subQuery, id).
			Limit(10).
			Find(&userList).Error
	}

	if err != nil {
		return UserList, err
	}

	return userList, nil
}
