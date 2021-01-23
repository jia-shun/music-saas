package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"music/config"
)

var (
	CONFIG config.Server
	LOG    *zap.Logger
	DB     *gorm.DB
	VIPER  *viper.Viper
	REDIS  *redis.Client
)
