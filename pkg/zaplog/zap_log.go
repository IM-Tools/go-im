/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package zaplog

import (
	"fmt"
	"go.uber.org/zap"
	"go_im/pkg/config"
	"go_im/pkg/helpler"
)


var (
	ZapLogger *zap.Logger
)

func InitZapLogger()  {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.zaplog", config.GetString("app.log_address"), helpler.GetNowFormatTodayTime()),
		"stdout",
	}
	// 创建logger实例
	logg, _ := cfg.Build()
	zap.ReplaceGlobals(logg) // 替换zap包中全局的logger实例
	ZapLogger = logg  // 注册到全局变量中
}

