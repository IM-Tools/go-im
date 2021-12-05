/**
  @author:panliang
  @data:2021/12/5
  @note
**/
package validates

import (
	"github.com/thedevsaddam/govalidator"
)

type EmailForm struct {
	Email string `valid:"email"`
}

func ValidateEmailForm(data EmailForm) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{"required", "email", "not_exists:im_users,email"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:email为必填项",
			"email:不是一个正确的邮箱格式",
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
