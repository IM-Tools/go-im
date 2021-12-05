/**
  @author:panliang
  @data:2021/6/21
  @note
**/
package user

import (
	"encoding/json"
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
	OauthType       int
	Status          int `json:"status"`
	OauthId         string
	CreatedAt       string `json:"created_at"`
	PasswordConfirm string ` gorm:"-" valid:"password_confirm"`
	LastLoginTime   string `json:"last_login_time"`
	Bio             string `json:"bio"`
	Sex             int    `json:"sex"`
	ClientType      int    `json:"client_type"`
	Age             int    `json:"age"`
}

type UsersWhiteList struct {
	ID            int64  `json:"id"`
	Email         string `valid:"email" json:"email"`
	Avatar        string `json:"avatar"`
	Name          string `json:"name" valid:"name"`
	OauthType     int
	Status        int `json:"status"`
	OauthId       string
	CreatedAt     string `json:"created_at"`
	Bio           string `json:"bio"`
	Sex           int    `json:"sex"`
	ClientType    int    `json:"client_type"`
	Age           int    `json:"Age"`
	LastLoginTime string `json:"last_login_time"`
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

func GetFriendListV2(user_id []int64) ([]Users, error) {
	var users []Users
	err := model.DB.Where("id in (?)", user_id).Find(&users).Error
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
