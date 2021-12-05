/**
  @author:panliang
  @data:2021/9/8
  @note
**/
package core

import (
	"github.com/gin-gonic/gin"
	"im_app/core/ws"
	conf "im_app/pkg/config"
	"im_app/pkg/zaplog"
	"im_app/router"
)

var appClusterModel = conf.GetBool("core.app_cluster_model")

func StartHttp() {

	app := gin.Default()

	SetupPool()

	go ws.ImManager.Start()

	if appClusterModel == true {
		go ws.StartRpc()
	}

	router.RegisterApiRoutes(app)
	router.RegisterIMRouters(app)

	app.Use(zaplog.Recover)

	_ = app.Run(":" + conf.GetString("core.port"))

}
