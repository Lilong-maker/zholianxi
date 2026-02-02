package init

import (
	"fmt"
	"log"
	"zholianxi/src/basic/config"
	"zholianxi/src/handler/model"

	"github.com/gospacex/gospacex/core/storage/conf"
	"github.com/gospacex/gospacex/core/storage/db/mysql"
)

func InitMysql() {
	config.DB, err = mysql.Init(true, "debug", conf.Cfg.Mysql)
	if err != nil {
		return
	}
	fmt.Println("数据库连接成功")
	log.Println("mysql连接成功")
	err = config.DB.AutoMigrate(&model.Goods{}, &model.Order{}, &model.User{})
	if err != nil {
		return
	}
	fmt.Println("表迁移成功")
}
