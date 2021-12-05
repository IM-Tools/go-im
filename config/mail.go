/**
  @author:panliang
  @data:2021/12/5
  @note
**/
package config

import "im_app/pkg/config"

func init() {
	config.Add("mail", config.StrMap{
		//weibo
		"driver": config.Env("MAIL_DRIVER", ""),

		"host": config.Env("MAIL_HOST", ""),
		"port": config.Env("MAIL_PORT", ""),
		"name": config.Env("MAIL_NAME", ""),
		//gitee
		"password":   config.Env("MAIL_PASSWORD", ""),
		"encryption": config.Env("MAIL_ENCRYPTION", ""),
		"address":    config.Env("MAIL_FROM_ADDRESS", ""),
		"from_name":  config.Env("MAIL_FROM_NAME", ""),
	})
}
