/**
  @author:panliang
  @data:2021/12/5
  @note
**/
package validates

import (
	"github.com/thedevsaddam/govalidator"
	"im_app/pkg/redis"
	"im_app/pkg/zaplog"
)

type UserRegisteredForm struct {
	Name            string `valid:"name"`
	Email           string `valid:"email"`
	Code            string `valid:"code"`
	Password        string `valid:"password"`
	PasswordConfirm string `valid:"password_confirm"`
}

func ValidateRegisteredForm(data UserRegisteredForm) map[string][]string {
	rules := govalidator.MapData{
		"name":             []string{"required", "between:3,20", "not_exists:im_users,name"},
		"email":            []string{"required", "email", "not_exists:im_users,email"},
		"code":             []string{"required", "min:4"},
		"password":         []string{"required", "between:6,20"},
		"password_confirm": []string{"required"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"between:用户名长度需在 3~20 之间",
		},
		"email": []string{
			"required:email为必填项",
			"email:不是一个正确的邮箱格式",
		},
		"code": []string{
			"required:code为必填项",
			"min:验证码长度为4",
		},
		"password": []string{
			"required:密码为必填项",
			"between:密码长度需在 6~20 之间",
		},
		"password_confirm": []string{
			"required:重复密码为必填项",
		},
	}
	// 3. 配置初始化
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	errs := govalidator.New(opts).ValidateStruct()
	if data.Password != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	StringCmd := redis.RedisDB.Get(data.Email)
	code := StringCmd.Val()
	zaplog.Info(data, code)
	if code != data.Code {
		errs["code"] = append(errs["code"], "邮箱验证码不正确！")
	}
	return errs
}
