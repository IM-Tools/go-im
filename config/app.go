/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package config

import (
	"go_im/pkg/config"
)

func init() {

	config.Add("app", config.StrMap{
		// 应用名称，暂时没有使用到
		"name": config.Env("APP_NAME", "GoIm"),
		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "production"),
		"port": config.Env("HTTP_PORT", "8000"),
		"jwt_secret": config.Env("APP_JWT_KEY", ""),
	})
}
