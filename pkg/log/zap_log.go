/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package log

import (
	"go.uber.org/zap"

)

var Logger *zap.Logger

func InitZapLogger()  {
	Logger, _ = zap.NewProduction()
}

