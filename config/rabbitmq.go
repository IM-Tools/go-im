/**
  @author:panliang
  @data:2021/9/15
  @note
**/
package config

import "im_app/pkg/config"

func init() {
	config.Add("rabbitmq", config.StrMap{
		"host":     config.Env("RABBITMQ_HOST", "localhost"),
		"port":     config.Env("RABBITMQ_PORT", "5672"),
		"user":     config.Env("RABBITMQ_USER", "guest"),
		"password": config.Env("RABBITMQ_PASSWORD", "guest"),
	})
}
