/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package config

import "go_im/pkg/config"

func init()  {
	config.Add("oauth", config.StrMap{
		//weibo
		"wb_client_id": config.Env("WEIBO_CLIENT_ID", ""),

		"wb_client_secret": config.Env("WEIBO_CLIENT_SECRET", ""),
		"wb_redirect_uri": config.Env("WEIBO_REDIRECT_URI", ""),
		//gitee
		"ge_client_id": config.Env("GITHUB_CLIENT_ID", ""),
		"ge_client_secret": config.Env("GITHUB_CALLBACK", ""),
		"ge_redirect_uri": config.Env("GITHUB_SECRET", ""),
	})
}
