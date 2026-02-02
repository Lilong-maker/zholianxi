package config

import (
	__ "zholianxi/src/basic/proto"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	GoodsClient __.GoodsClient
	DB          *gorm.DB
	Rdb         *redis.Client
)
