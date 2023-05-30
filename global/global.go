package global

import (
	"gin_web_server/config"
	"go.uber.org/zap"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

var (
	VIPER *viper.Viper

	GvaConfig config.Server

	GvaLog *zap.Logger

	BlackCache local_cache.Cache

	GlaDb *sqlx.DB
)
