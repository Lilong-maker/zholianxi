package init

import (
	"fmt"
	"zholianxi/bff/basic/config"

	"github.com/gospacex/gospacex/core/storage/cache/redis"
	"github.com/gospacex/gospacex/core/storage/conf"
)

func InitRedis() {
	err = redis.Init(false, conf.Cfg.Redis)
	if err != nil {
		return
	}
	config.Rdb = redis.RC
	fmt.Println("redis连接成功")
}
