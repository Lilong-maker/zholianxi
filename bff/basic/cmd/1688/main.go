package main

import (
	"fmt"

	"gitee.com/zuyanlongnb666/crawl/Upload"
	"gitee.com/zuyanlongnb666/crawl/cra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	DB  *gorm.DB
)

func main() {
	Mysql()
	crawl := cra.Crawl(
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36",
		"https://sale.1688.com/factory/u0vjcc4j.html?spm=a260k.home2025.centralDoor.ddoor.66333597BBbHgE&topOfferIds=1005591171200")
	for _, m := range crawl {

		goods := CrawlGoods{
			Title: m["title"],
			Price: m["price"],
			Img:   m["src"],
		}
		Upload.Upload(m["src"], "111")
		err = goods.CreateAdd(DB)
		if err != nil {
			return
		}
		fmt.Println("添加成功")
	}
}

func Mysql() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:4ay1nkal3u8ed77y@tcp(115.190.43.83:3306)/p2308a?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("连接成功")
	err = DB.AutoMigrate(CrawlGoods{})
	if err != nil {
		return
	}
	fmt.Println("迁移成功")
}

type CrawlGoods struct {
	Title string `gorm:"type:varchar(50)"`
	Price string `gorm:"type:varchar(50)"`
	Img   string `gorm:"type:varchar(500)"`
}

func (g *CrawlGoods) CreateAdd(d *gorm.DB) error {
	return d.Debug().Create(&g).Error
}
