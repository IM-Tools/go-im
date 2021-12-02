/**
  @author:panliang
  @data:2021/6/22
  @note
**/
package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"im_app/core/http/models/user"
	NewJwt "im_app/pkg/jwt"
	"im_app/pkg/response"
)

var (
	token  string
	err    error
	claims *NewJwt.CustomClaims
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token = c.DefaultQuery("token", c.GetHeader("authorization"))
		err, token = ValidatedToken(token)
		if err != nil {
			response.ErrorResponse(401, err.Error()).WriteTo(c)
			c.Abort()
			return
		}
		jwt := NewJwt.NewJWT()
		claims, err = jwt.ParseToken(token)
		if err != nil {
			response.ErrorResponse(401, err.Error()).WriteTo(c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		setAuthUser(c)

		c.Next()
	}
}

// ValidateToken 验证token
func ValidatedToken(token string) (error, string) {
	if len(token) == 0 {
		return errors.New("Token 不能为空"), ""
	}

	t := strings.Split(token, "Bearer ")
	if len(t) > 1 {
		return nil, t[1]
	}
	return nil, token
}

// setAuthUser 设置登录用户
func setAuthUser(c *gin.Context) {
	claims = c.MustGet("claims").(*NewJwt.CustomClaims)
	user.AuthUser = &user.Users{
		ID:     claims.ID,
		Email:  claims.Email,
		Avatar: claims.Avatar,
		Name:   claims.Name,
	}
}
