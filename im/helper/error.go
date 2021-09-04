/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package helper

import "go.uber.org/zap"

// records  error log
func Error(msg string ,err error)  {
	logger, _ := zap.NewProduction()
	logger.Error(msg,zap.Error(err))
}
