package init

import (
	"fmt"

	"github.com/gospacex/gospacex/core/storage/conf"
)

func InitViper() {
	conf.ParseConfig("../../../")
	fmt.Println("配置文件加载成功")
}
