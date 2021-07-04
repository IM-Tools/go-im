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
	userModel "go_im/bin/http/models/user"
	"go_im/bin/http/validates"
	"go_im/bin/oauth"
	"go_im/pkg/config"
	"go_im/pkg/helpler"
	"go_im/pkg/jwt"
	"go_im/pkg/model"
	"go_im/pkg/response"
	"strconv"
	"time"
)

//定义一个结构体 用户该方法的引用
type AuthController struct{}
type WeiBoController struct{}

//定义结构数据格式
type Me struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	Avatar         string `json:"avatar"`
	Email          string `json:"email"`
	Token          string `json:"token"`
	ExpirationTime int64  `json:"expiration_time"`
}

func (*AuthController) GiteeCallBack(c *gin.Context) {
	code := c.Query("code")
	if len(code) == 0 {
		response.FailResponse(403, "参数不正确~").ToJson(c)
	}
	//微博授权
	//access_token := oauth.GetGiteeAccessToken(&code)

	//UserInfo := oauth.GetGiteeUserInfo(&access_token)

}

//登录并返回用户信息与token

func (*AuthController) Me(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	data := map[string]interface{}{
		"id":     claims.ID,
		"name":   claims.Name,
		"avatar": claims.Avatar,
		"email":  claims.Email,
	}
	response.SuccessResponse(data, 200).ToJson(c)
}

func (that *AuthController) Login(c *gin.Context) {
	var params validates.LoginParams
	var users userModel.Users
	_ = c.ShouldBind(&params)

	model.DB.Model(&userModel.Users{}).Where("name = ?", params.Name).Find(&users)
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

//重新刷新用户信息
func (*AuthController) Refresh(c *gin.Context) {
	//claims := c.MustGet("claims").(*jwt.CustomClaims)
	//
	//sign_key            := config.GetString("app.jwt.sign_key")
	//expiration_time     := config.GetInt("app.jwt.expiration_time")
	//
	//user := userModel.Users{}
	//
	//result := model.DB.Where("id =?",claims.ID).First(&user)
	//
	//if result.Error != nil {
	//	c.JSON(http.StatusOK, map[string]interface{}{
	//		"code": 500,
	//		"msg":  "用户信息不存在",
	//	})
	//	return
	//}
	//
	//j :=jwt.JWT{
	//	[]byte(sign_key),
	//}
	//claimsString := jwt.CustomClaims{strconv.FormatInt(user.ID,10),
	//	user.Name,
	//	user.Avatar,
	//	user.Email,
	//	jwtgo.StandardClaims{
	//		NotBefore: time.Now().Unix() - 1000,
	//		ExpiresAt: time.Now().Unix() + int64(expiration_time),
	//		Issuer: sign_key,
	//	}}
	//token,err := j.RefreshToken(claimsString)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//data := map[string]interface{}{
	//	"code":200,
	//	"msg":"登录成功",
	//	"data": map[string]interface{}{
	//		"token":token,
	//		"id":user.ID,
	//		"name":user.Name,
	//		"avatar":user.Avatar,
	//		"email":user.Email,
	//		"expiration_time":expiration_time,
	//	},
	//}
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"code":200,
	//	"msg":"登录成功",
	//	"data": data,
	//})

}

//微博授权接口
func (*WeiBoController) WeiBoCallBack(c *gin.Context) {
	code := c.Query("code")
	if len(code) == 0 {
		response.FailResponse(403, "参数不正确~").ToJson(c)
	}
	//微博授权
	access_token := oauth.GetWeiBoAccessToken(&code)
	UserInfo := oauth.GetWeiBoUserInfo(&access_token)

	users := userModel.Users{}
	//oauth_id := gjson.Get(UserInfo,"id").Raw
	oauth_id := "5878370732"

	isThere := model.DB.Where("oauth_id = ?", oauth_id).First(&users)
	fmt.Println(users)
	//用户未授权
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
			response.FailResponse(500, "用户微博授权失败").ToJson(c)
		} else {
			generateToken(c, &userData)
		}
	} else {
		generateToken(c, &users)
	}
}

// 给用户颁发token
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
