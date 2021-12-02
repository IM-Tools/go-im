/**
  @author:panliang
  @data:2021/8/18
  @note
**/
package pool

import (
	"github.com/panjf2000/ants/v2"
	"im_app/pkg/config"
)

var AntsPool *ants.Pool

func ConnectPool() *ants.Pool {
	//设置数量
	AntsPool, _ = ants.NewPool(config.GetInt("core.go_coroutines"))
	return AntsPool
}
