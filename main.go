package main

import (
	"gin_web_server/core"
	"gin_web_server/global"
	"gin_web_server/initialize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const AppMode = "debug" // 运行环境，主要有三种：debug、test、release
func main() {
	gin.SetMode(AppMode)
	// TODO：1.配置初始化
	global.VIPER = core.InitializeViper() // 初始化Viper
	// TODO：2.日志
	global.GvaLog = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GvaLog)

	// TODO：3.数据库连接
	global.GlaDb = initialize.SqlDB()
	// TODO：4.其他初始化
	// TODO：5.启动服务
	core.RunWindowsServer()
}
