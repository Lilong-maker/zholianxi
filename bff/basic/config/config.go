package config

import (
	__ "zholianxi/src/basic/proto"

	"gorm.io/gorm"
)

var (
	GoodsClient __.GoodsClient
	DB          *gorm.DB
)
