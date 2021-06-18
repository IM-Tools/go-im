/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package config

import "go_im/pkg/config"

func init()  {
	config.Add("oauth", config.StrMap{
		// 应用名称，暂时没有使用到
		"client_id": config.Env("WEIBO_CLIENT_ID", ""),
		// 当前环境，用以区分多环境
		"client_secret": config.Env("WEIBO_CLIENT_SECRET", ""),
		"redirect_uri": config.Env("WEIBO_REDIRECT_URI", ""),
	})
}
