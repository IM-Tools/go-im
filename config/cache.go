/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package config

import "go_im/pkg/config"

func init() {
	config.Add("cache", config.StrMap{
		"redis": map[string]interface{}{
			// redis连接信息
			"addr":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"password": config.Env("REDIS_PASSWORD", ""),
			"db":       config.Env("REDIS_DB", 0),
		},
	})
}
