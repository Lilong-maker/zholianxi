package init

import (
	"fmt"
	"zholianxi/src/basic/config"

	"github.com/olivere/elastic/v7"
)

var err error

func InitEs() {
	config.Elastic, err = elastic.NewClient(
		elastic.SetURL("http://115.190.43.83:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("es连接成功")
}
