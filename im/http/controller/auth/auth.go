/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package auth

import (
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	userModel "go_im/im/http/models/user"
	"go_im/im/http/validates"
	"go_im/im/utils"
	"go_im/pkg/config"
	"go_im/pkg/helpler"
	"go_im/pkg/jwt"
	"go_im/pkg/model"
	"go_im/pkg/response"
	"strconv"
	"time"
)
type (
	AuthController struct{}
	WeiBoController struct{}
	Me struct {
		ID             uint64 `json:"id"`
		Name           string `json:"name"`
		Avatar         string `json:"avatar"`
		Email          string `json:"email"`
		Token          string `json:"token"`
		ExpirationTime int64  `json:"expiration_time"`
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

// @Summary 这是一个登录接口
// @Description 登录接口
// @Tags 登录接口
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "账号"
// @Param password formData string true "密码"
// @Success 200
// @Router /login [post]
func (that *AuthController) Login(c *gin.Context) {
	_user := userModel.Users{
		Name: c.PostForm("name"),
		Password: c.PostForm("password"),
	}
	errs := validates.ValidateLoginForm(_user)
	if len(errs) >0 {
		response.ErrorResponse(500,"参数错误",errs).WriteTo(c)
		return
	}
	var users userModel.Users
	model.DB.Model(&userModel.Users{}).Where("name = ?", _user.Name).Find(&users)
	if users.ID == 0 {
		response.FailResponse(403, "用户不存在").ToJson(c)
		return
	}
	if !helpler.ComparePasswords(users.Password, _user.Password) {
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
	access_token := utils.GetWeiBoAccessToken(&code)
	UserInfo := utils.GetWeiBoUserInfo(&access_token)
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
			response.FailResponse(500, "用户微博授权失败").ToJson(c)
		} else {
			generateToken(c, &userData)
		}
	} else {
		generateToken(c, &users)
	}
}
//d55894389fe5bca5833cf28aa4384b04 EngV58oPYmhoIKIbBkFh3KLcRMqr9ke8x1I1KcSONwn
func (*AuthController) WxCallback(c *gin.Context)  {
	response.SuccessResponse().ToJson(c)
	return
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
