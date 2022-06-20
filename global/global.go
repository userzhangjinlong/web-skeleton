package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"linkr-frame/config"
)

//统一管理全局常量
var (
	Config   *config.System
	LinkItDB *gorm.DB
	SchemaDB *gorm.DB
	Redis    *redis.Client
)
