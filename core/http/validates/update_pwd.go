/**
  @author:panliang
  @data:2021/12/7
  @note
**/
package validates

import "github.com/thedevsaddam/govalidator"

type PwdFrom struct {
	Password        string `valid:"password"`
	PasswordConfirm string `valid:"password_confirm"`
	NewPassword     string `valid:"new_password"`
}

func ValidatePwdFrom(data PwdFrom) map[string][]string {
	rules := govalidator.MapData{
		"new_password":     []string{"required", "between:6,20"},
		"password":         []string{"required", "between:6,20"},
		"password_confirm": []string{"required"},
	}
	messages := govalidator.MapData{

		"new_password": []string{
			"required:新密码为必填项",
			"between:新密码长度需在 6~20 之间",
		},
		"password": []string{
			"required:旧密码为必填项",
			"between:旧密码长度需在 6~20 之间",
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
	if data.NewPassword != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	if data.NewPassword == data.Password {
		errs["password"] = append(errs["password"], "新密码和旧密码一至！")
	}
	return errs
}
