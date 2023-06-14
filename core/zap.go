package core

import (
	"fmt"
	"gin_web_server/core/internal"
	"gin_web_server/global"
	"gin_web_server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GvaConfig.Zap.Director); !ok { // 判断是否有Director文件夹

		fmt.Printf("create %v directory\n", global.GvaConfig.Zap.Director)
		_ = os.Mkdir(global.GvaConfig.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))
	log.Println()
	if global.GvaConfig.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}
