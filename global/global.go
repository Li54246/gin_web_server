package global

import (
	"gin_web_server/config"
	"go.uber.org/zap"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
)

var (
	VIPER *viper.Viper

	GVA_CONFIG config.Server

	GVA_LOG *zap.Logger

	BlackCache local_cache.Cache
)
