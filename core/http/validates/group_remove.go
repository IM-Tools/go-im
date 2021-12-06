/**
  @author:panliang
  @data:2021/12/6
  @note
**/
package validates

import (
	"github.com/thedevsaddam/govalidator"
)

type RemoveUserFormGroupFrom struct {
	GroupId string `valid:"group_id"`
	UserId  string `valid:"user_id"`
}

func ValidateRemoveGroupForm(data RemoveUserFormGroupFrom) map[string][]string {
	rules := govalidator.MapData{
		"group_id": []string{"required", "between:3,20", "exists:im_groups,id"},
		"user_id":  []string{"required"},
	}
	messages := govalidator.MapData{
		"group_id": []string{
			"required:群聊id必填项",
			"exists:群聊不存在",
		},
		"user_id": []string{
			"required:用户id为必填项",
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
