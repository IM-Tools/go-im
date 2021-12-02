/**
  @author:panliang
  @data:2021/9/7
  @note
**/
package validates

import (
	"github.com/thedevsaddam/govalidator"

	user2 "im_app/core/http/models/user"
)

func ValidateLoginForm(data user2.Users) map[string][]string {
	rules := govalidator.MapData{
		"name":     []string{"required"},
		"password": []string{"required", "min:5"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:长度需大于 6",
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
	return errs
}
