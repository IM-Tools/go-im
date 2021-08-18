/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package auth

import (
	"fmt"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	userModel "go_im/im/http/models/user"
	"go_im/im/http/validates"
	"go_im/im/oauth"
	"go_im/pkg/config"
	"go_im/pkg/helpler"
	"go_im/pkg/jwt"
	log2 "go_im/pkg/log"
	"go_im/pkg/model"
	"go_im/pkg/response"
	"strconv"
	"time"
)

type AuthController struct{}
type WeiBoController struct{}

type Me struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	Avatar         string `json:"avatar"`
	Email          string `json:"email"`
	Token          string `json:"token"`
	ExpirationTime int64  `json:"expiration_time"`
}

func (*AuthController) Me(c *gin.Context) {
	user := userModel.AuthUser
	response.SuccessResponse(user, 200).ToJson(c)
}

func (that *AuthController) Login(c *gin.Context) {
	var params validates.LoginParams
	var users userModel.Users
	_ = c.ShouldBind(&params)

	model.DB.Model(&userModel.Users{}).Where("name = ?", params.Name).Find(&users)
	fmt.Println(users)
	if users.ID == 0 {
		response.FailResponse(403, "用户不存在").ToJson(c)
		return
	}
	if !helpler.ComparePasswords(users.Password, params.Password) {
		response.FailResponse(403, "账号或者密码错误").ToJson(c)
		return
	}
	generateToken(c, &users)
}

func (*WeiBoController) WeiBoCallBack(c *gin.Context) {
	code := c.Query("code")
	if len(code) == 0 {
		response.FailResponse(403, "参数不正确~").ToJson(c)
	}
	access_token := oauth.GetWeiBoAccessToken(&code)
	UserInfo := oauth.GetWeiBoUserInfo(&access_token)
	users := userModel.Users{}
	oauth_id := gjson.Get(UserInfo,"id").Raw
	isThere := model.DB.Where("oauth_id = ?", oauth_id).First(&users)
	if isThere.Error != nil {
		userData := userModel.Users{
			Email:           gjson.Get(UserInfo, "email").Str,
			Password:        helpler.HashAndSalt("123456"),
			PasswordComfirm: helpler.HashAndSalt("123456"),
			OauthId:         gjson.Get(UserInfo, "id").Raw,
			Avatar:          gjson.Get(UserInfo, "avatar_large").Str,
			Name:            gjson.Get(UserInfo, "name").Str,
			OauthType:       1,
			CreatedAt:       time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		}
		result := model.DB.Create(&userData)

		if result.Error != nil {
			log2.Warning(result.Error.Error())
			response.FailResponse(500, "用户微博授权失败").ToJson(c)
		} else {
			generateToken(c, &userData)
		}
	} else {
		generateToken(c, &users)
	}
}


func generateToken(c *gin.Context, user *userModel.Users) {
	sign_key := config.GetString("app.jwt.sign_key")
	expiration_time := config.GetInt64("app.jwt.expiration_time")

	j := &jwt.JWT{
		[]byte(sign_key),
	}
	claims := jwt.CustomClaims{strconv.FormatUint(user.ID, 10),
		user.Name,
		user.Avatar,
		user.Email,
		jwtGo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + expiration_time,
			Issuer:    sign_key,
		}}
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailResponse(403, "jwt token颁发失败~").ToJson(c)
		return
	} else {
		data := new(Me)
		data.ID = user.ID
		data.Name = user.Name
		data.Avatar = user.Avatar
		data.Email = user.Email
		data.Token = token
		data.ExpirationTime = expiration_time
		response.SuccessResponse(data, 200).ToJson(c)
		return
	}
}
