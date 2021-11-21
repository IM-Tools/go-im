/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"im_app/pkg/config"
	"im_app/pkg/zaplog"
	"strconv"
	"time"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = config.GetString("app.jwt.sign_key")
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID           string `json:"userId"`
	Name, Avatar string
	Email        string `valid:"email"`
	ClientType int `json:"client_type"` //用户设备
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey

}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey

}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func GenerateToken(uid uint64,Name string,avatar string,email string,ClientType int) (token string)  {
	sign_key := config.GetString("app.jwt.sign_key")
	expiration_time := config.GetInt64("app.jwt.expiration_time")
	j := &JWT{
		[]byte(sign_key),
	}
	claims := CustomClaims{strconv.FormatUint(uid, 10),
		Name,
		avatar,
		email,
		ClientType,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + expiration_time,
			Issuer:    sign_key,
		}}
	token,err := j.CreateToken(claims)

	if err != nil {
		zaplog.Error("----token颁发失败",err)
	}

	return token

}


// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

//更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
