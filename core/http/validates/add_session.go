/**
  @author:panliang
  @data:2021/12/21
  @note
**/
package validates

import (
	"github.com/thedevsaddam/govalidator"
	"im_app/core/http/models/user"
	"im_app/pkg/model"
)

type AddSessionFrom struct {
	UserId      string `valid:"user_id"`
	ChannelType string `valid:"channel_type"`
}

func ValidateAddSession(data AddSessionFrom) map[string][]string {
	rules := govalidator.MapData{
		"user_id":      []string{"required"},
		"channel_type": []string{"required"},
	}
	messages := govalidator.MapData{
		"user_id": []string{
			"required:好友id或群聊id不能为空",
		},
		"channel_type": []string{
			"required:好友id为必填项",
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
	var count int64
	model.DB.Table("im_sessions").
		Where("m_id=? and f_id=? and status=0 and channel_type=?", user.AuthUser.ID, data.UserId, data.ChannelType).
		Count(&count)
	if count != 0 {
		errs["code"] = append(errs["m_id"], "会话已存在")
		return errs
	}
	return errs

}
