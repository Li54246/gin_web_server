package core

import (
	"fmt"
	"gin_web_server/global"
	"gin_web_server/initialize"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	//r := gin.Default()
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong")
	//	global.GVA_LOG.Error("获取失败!")
	//})
	//
	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.App.Port)
	//global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	//
	//fmt.Println(`address`, address)
	//err := r.Run(address)
	//if err != nil {
	//	return
	//}

	//if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
	//	// 初始化redis服务
	//	initialize.Redis()
	//}
	//
	//// 从db加载jwt数据
	//if global.GVA_DB != nil {
	//	system.LoadAll()
	//}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GvaConfig.App.Port)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GvaLog.Info("server run success on ", zap.String("address", address))

	fmt.Println(`
	启动成功
`, address)
	global.GvaLog.Error(s.ListenAndServe().Error())
}
