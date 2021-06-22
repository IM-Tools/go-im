/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package auth

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	userModel "go_im/bin/http/models/user"
	"go_im/bin/oauth"
	"go_im/pkg/config"
	"go_im/pkg/helpler"
	"go_im/pkg/jwt"
	"strconv"
	"time"

	"go_im/pkg/model"
	"net/http"
)

//定义一个结构体 用户该方法的引用
type AuthController struct{}
type WeiBo struct{}

//登录并返回用户信息与token
func (*AuthController)Me(c *gin.Context)  {
	claims := c.MustGet("claims").(*jwt.CustomClaims)

	c.JSON(http.StatusForbidden,map[string]interface{}{
		"code":200,
		"msg":"success",
		"data": map[string]interface{}{
			"id":claims.ID,
			"name":claims.Name,
			"avatar":claims.Avatar,
			"email":claims.Email,
		},
	})

}

//微博授权接口
func (*WeiBo)WeiBoCallBack (c *gin.Context)  {
	code := c.Query("code")
	if len(code) ==0 {
		c.JSON(http.StatusForbidden,map[string]interface{}{
			"msg": "参数不正确～",
			"code":code,
		})
	}
	//微博授权
	access_token := oauth.GetAccessToken(&code)
	UserInfo := oauth.GetUserInfo(&access_token)

	users :=userModel.Users{}

	isThere := model.DB.Where("oauth_id = ?",gjson.Get(UserInfo,"id").Raw).First(&users)
	//用户未授权
	if isThere.Error != nil {

		userData := userModel.Users{
			Email:gjson.Get(UserInfo,"email").Str,
			Password:helpler.HashAndSalt("123456"),
			PasswordComfirm:  helpler.HashAndSalt("123456"),
			OauthId: gjson.Get(UserInfo,"id").Raw,
			Avatar: gjson.Get(UserInfo,"avatar_large").Str,
			Name: gjson.Get(UserInfo,"name").Str,
			OauthType:1,
			CreatedAt:time.Unix(time.Now().Unix(), 0,).Format("2006-01-02 15:04:05"),
		}
		result := model.DB.Create(&userData)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code": 500,
				"msg":  "用户微博授权失败",
			})
		} else {
			generateToken(c,&userData)
		}
	} else {

		generateToken(c,&users)

	}
}


// 给用户颁发token
func generateToken(c *gin.Context,user *userModel.Users) {

	fmt.Println("测试users",user)
	sign_key            := config.GetString("app.jwt.sign_key")
	expiration_time     := config.GetInt("app.jwt.expiration_time")

	fmt.Println(expiration_time)

	fmt.Println(user.ID)

	j :=&jwt.JWT{
		[]byte(sign_key),
	}
	claims := jwt.CustomClaims{strconv.FormatInt(user.ID,10),
		user.Name,
		user.Avatar,
		user.Email,
		jwtgo.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,
		ExpiresAt: time.Now().Unix() + int64(expiration_time),
		Issuer: sign_key,
	}}

	token,err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusForbidden, map[string]interface{}{
			"code":403,
			"msg":"jwt token颁发失败--",
		})
		return
	} else {
		data := map[string]interface{}{
			"code":200,
			"msg":"登录成功",
			"data": map[string]interface{}{
				"token":token,
				"id":user.ID,
				"name":user.Name,
				"avatar":user.Avatar,
				"email":user.Email,
				"expiration_time":expiration_time,
			},
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":200,
			"msg":"登录成功",
			"data": data,
		})
		return
	}
}
