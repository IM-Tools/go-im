/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"im_app/core/http/models/friend"
	userModel "im_app/core/http/models/user"
	"im_app/core/http/services"
	"im_app/core/http/validates"
	"im_app/core/utils"
	"im_app/core/ws"
	"im_app/pkg/config"
	"im_app/pkg/helpler"
	"im_app/pkg/jwt"
	"im_app/pkg/model"
	"im_app/pkg/redis"
	"im_app/pkg/response"
	"strconv"
	"time"
)

type (
	AuthController  struct{}
	WeiBoController struct{}
	Me              struct {
		ID             int64  `json:"id"`
		Name           string `json:"name"`
		Avatar         string `json:"avatar"`
		Email          string `json:"email"`
		Token          string `json:"token"`
		ExpirationTime int64  `json:"expiration_time"`
		Sex            int    `json:"sex"`
		ClientType     int    `json:"client_type"`
		Bio            string `json:"bio"`
	}
)

// @BasePath /api

// @Summary 获取用户信息接口
// @Description 获取用户信息接口
// @Tags 获取用户信息接口
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Produce json
// @Success 200
// @Router /me [post]
func (*AuthController) Me(c *gin.Context) {
	user := userModel.AuthUser
	response.SuccessResponse(user, 200).ToJson(c)
}

// @BasePath /api

// @Summary 更新用户数据
// @Description 更新用户数据
// @Tags 更新用户数据
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param bio formData string false "个性签名"
// @Param six formData int false "性别"

// @Produce json
// @Success 200
// @Router /Update [put]
func (*AuthController) Update(c *gin.Context) {
	bio := c.PostForm("bio")
	sex := c.PostForm("sex")
	user := userModel.AuthUser
	var users userModel.Users
	if len(sex) != 0 {
		sex, _ := strconv.Atoi(sex)
		users.Sex = sex
	}
	users.Bio = bio

	model.DB.Table("users").Where("id", user.ID).First(&users)
	users.Bio = bio
	model.DB.Save(&users)
	response.SuccessResponse().WriteTo(c)
	return
}

// @BasePath /api

// @Summary 这是一个登录接口
// @Description 登录接口
// @Tags 登录接口
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "账号"
// @Param password formData string true "密码"
// @Param client_type formData string false "客户端类型 0.网页端登录 1.设备端登录"
// @Success 200
// @Router /login [post]
func (that *AuthController) Login(c *gin.Context) {
	_user := userModel.Users{
		Name:     c.PostForm("name"),
		Password: c.PostForm("password"),
	}

	ClientType, _ := strconv.Atoi(c.DefaultPostForm("client_type", "0"))

	errs := validates.ValidateLoginForm(_user)
	if len(errs) > 0 {
		response.ErrorResponse(500, "参数错误", errs).WriteTo(c)
		return
	}
	var users userModel.Users
	model.DB.Model(&userModel.Users{}).
		Where("name = ? or email = ?", _user.Name, _user.Name).
		Find(&users)
	if users.ID == 0 {
		response.FailResponse(403, "用户不存在").ToJson(c)
		return
	}
	if !helpler.ComparePasswords(users.Password, _user.Password) {
		response.FailResponse(403, "账号或者密码错误").ToJson(c)
		return
	}

	users.ClientType = ClientType
	model.DB.Model(&userModel.Users{}).Save(&users)
	// 挤下线操作
	if users.Status == 1 {
		ws.CrowdedOffline(int64(users.ID))
	}
	token := jwt.GenerateToken(users.ID, users.Name, users.Avatar, users.Email, ClientType)
	data := getMe(token, &users)
	response.SuccessResponse(data, 200).ToJson(c)
}

func (*WeiBoController) WeiBoCallBack(c *gin.Context) {
	code := c.Query("code")
	if len(code) == 0 {
		response.FailResponse(403, "参数不正确~").ToJson(c)
	}
	access_token := utils.GetWeiBoAccessToken(&code)
	UserInfo := utils.GetWeiBoUserInfo(&access_token)
	users := userModel.Users{}
	oauth_id := gjson.Get(UserInfo, "id").Raw
	isThere := model.DB.Where("oauth_id = ?", oauth_id).First(&users)
	if isThere.Error != nil {
		userData := userModel.Users{
			Email:           gjson.Get(UserInfo, "email").Str,
			Password:        helpler.HashAndSalt("123456"),
			PasswordConfirm: helpler.HashAndSalt("123456"),
			OauthId:         gjson.Get(UserInfo, "id").Raw,
			Avatar:          gjson.Get(UserInfo, "avatar_large").Str,
			Name:            gjson.Get(UserInfo, "name").Str,
			OauthType:       1,
			CreatedAt:       time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
			LastLoginTime:   time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		}
		result := model.DB.Create(&userData)

		if result.Error != nil {
			response.FailResponse(500, "用户微博授权失败").ToJson(c)
		} else {
			// 执行默认添加好友逻辑
			friend.AddDefaultFriend(userData.ID)
			token := jwt.GenerateToken(userData.ID, userData.Name, userData.Avatar, userData.Email, 0)
			data := getMe(token, &userData)
			response.SuccessResponse(data, 200).ToJson(c)
		}
	} else {
		token := jwt.GenerateToken(users.ID, users.Name, users.Avatar, users.Email, 0)
		data := getMe(token, &users)
		response.SuccessResponse(data, 200).ToJson(c)

	}
}

// @BasePath /api

// @Summary 发送注册邮箱验证码
// @Description 发送注册邮箱验证码接口
// @Tags 发送注册邮箱验证码接口
// @Param email query string false "邮箱"
// @Success 200
// @Router /seedRegisteredEmail [get]
func (*AuthController) SeedRegisteredEmail(c *gin.Context) {

	_email := validates.EmailForm{
		c.Query("email"),
	}
	errs := validates.ValidateEmailForm(_email)

	if len(errs) > 0 {
		response.FailResponse(500, "error", errs).ToJson(c)
		return
	}

	code := helpler.CreateEmailCode()
	emailService := new(services.EmailService)
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>IM注册邮件</title>
</head>
<style>
    body {

    }
    .mail{
        margin: 0 auto;
    }
    span{
        color: #fff;
        font-weight: bold;
        font-size: 15px;
        background-color: #c56e57;
        padding: 2px;
    }
</style>
<body>
<div class="mail">
    <h3>您好:您正在注册im应用账号!</h3>
    <p>下面是您的验证码:<span>%s</span>请注意查收!谢谢</p>
    <h3>如果可以请给项目点个star,开源不易,你的star就是对我们最大的认可～<a target="_blank" href="https://github.com/pl1998/go-im">https://github.com/pl1998/go-im</a> </h3>
</div>
</body>
</html>`, code)

	err := emailService.SendEmail(_email.Email, "欢迎👏注册IM账号,这是一封邮箱验证码的邮件!🎉🎉🎉", html)
	if err != nil {
		response.FailResponse(500, "邮件发送失败,请检查是否是可用邮箱").ToJson(c)
		return
	}
	redis.RedisDB.Set(_email.Email, code, time.Minute*5)

	response.SuccessResponse().ToJson(c)
	return
}

// @BasePath /api

// @Summary 注册用户
// @Description 注册用户接口
// @Tags 注册用户
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "用户名"
// @Param email formData string true "邮箱"
// @Param password formData string true "密码"
// @Param password_confirm formData string true "确认密码"
// @Param code formData string true "验证码"
// @Success 200
// @Router /registered [post]
func (*AuthController) Registered(c *gin.Context) {

	_user := validates.UserRegisteredForm{
		Name:            c.PostForm("name"),
		Email:           c.PostForm("email"),
		Password:        c.PostForm("password"),
		Code:            c.PostForm("code"),
		PasswordConfirm: c.PostForm("password_confirm"),
	}
	errs := validates.ValidateRegisteredForm(_user)

	if len(errs) > 0 {
		response.FailResponse(500, "error", errs).ToJson(c)
		return
	}

	// 注册用户信息
	userData := userModel.Users{
		Email:           _user.Email,
		Password:        helpler.HashAndSalt(_user.Password),
		PasswordConfirm: helpler.HashAndSalt(_user.Password),
		Name:            _user.Name,
		CreatedAt:       time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		LastLoginTime:   time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
	}

	result := model.DB.Create(&userData)

	if result.Error != nil {
		response.FailResponse(500, "用户账号注册失败,请联系管理员").ToJson(c)
		return
	}
	//添加好友逻辑
	friend.AddDefaultFriend(userData.ID)

	response.SuccessResponse().ToJson(c)

}

// @BasePath /api

// @Summary 绑定用户邮箱
// @Description 绑定用户邮箱接口
// @Tags 绑定用户邮箱
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "邮箱"
// @Success 200
// @Router /bindingEmail [post]
func (*AuthController) BindingEmail(c *gin.Context) {

	_email := validates.EmailForm{
		c.PostForm("email"),
	}
	errs := validates.ValidateEmailForm(_email)

	if len(errs) > 0 {
		response.FailResponse(500, "error", errs).ToJson(c)
		return
	}

	user := userModel.AuthUser

	user.Email = _email.Email

	model.DB.Table("im_users").Where("id=?", user.ID).Update("email", _email.Email)

	response.SuccessResponse().ToJson(c)

	return

}

func (*AuthController) WxCallback(c *gin.Context) {
	response.SuccessResponse().ToJson(c)
	return
}

func getMe(token string, user *userModel.Users) *Me {
	expiration_time := config.GetInt64("core.jwt.expiration_time")
	data := new(Me)
	data.ID = user.ID
	data.Name = user.Name
	data.Avatar = user.Avatar
	data.Email = user.Email
	data.Token = token
	data.ExpirationTime = expiration_time
	data.Bio = user.Bio
	data.Sex = user.Sex
	data.ClientType = user.ClientType
	return data
}