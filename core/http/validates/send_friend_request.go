/**
  @author:panliang
  @data:2021/12/14
  @note
**/
package validates

import (
	"github.com/thedevsaddam/govalidator"
	"im_app/pkg/model"
)

type SendFriendRequestFrom struct {
	FId string `valid:"f_id"`
	MId string `valid:"m_id"`
}

func ValidateSendFriendRequestFrom(data SendFriendRequestFrom) map[string][]string {

	rules := govalidator.MapData{
		"f_id": []string{"required"},
		"m_id": []string{"required"},
	}

	messages := govalidator.MapData{
		"f_id": []string{
			"required:好友id为必填项",
		},
		"m_id": []string{
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
	var count int64
	model.DB.Table("im_friends").Where("m_id=? and f_id=?", data.FId, data.MId).Count(&count)
	if count != 0 {
		errs["code"] = append(errs["m_id"], "已经是好友了！请勿重复添加")
		return errs
	}
	var send_count int64
	model.DB.Table("im_friends").Where("m_id=? and f_id=? and status=0", data.FId, data.MId).Count(&count)
	if send_count != 0 {
		errs["code"] = append(errs["m_id"], "等待好友接受！请勿重复申请")
		return errs
	}
	return errs
}
