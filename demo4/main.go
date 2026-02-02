package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	//创建浏览器
	url := launcher.New().Headless(true).
		Set("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36").
		MustLaunch()
	//建立连接
	browser := rod.New().ControlURL(url).MustConnect()
	//延迟关闭
	//defer browser.Close()
	listPage := browser.MustPage("https://sale.1688.com/factory/u0vjcc4j.html?spm=a260k.home2025.centralDoor.ddoor.66333597BBbHgE&topOfferIds=1005591171200")

	listPage.MustWaitLoad()

	err := listPage.Timeout(20 * time.Second).MustElement(".offerItem").WaitVisible()
	if err != nil {
		fmt.Println("加载商品失败")
		return
	}
	elements := listPage.MustElements(".offerItem")
	for i, item := range elements {
		priceSpans := item.MustElements("span.text")
		var priceParts []string
		for _, span := range priceSpans {
			text := span.MustText()
			if strings.Contains(text, "¥") || strings.Contains(text, ".") ||
				(len(text) > 0 && text[0] >= '0' && text[0] <= '9') {
				priceParts = append(priceParts, text)
			}
		}
		priceStr := strings.Join(priceParts, "")
		fmt.Printf("%d. 价格: %s\n", i+1, priceStr)
	}
}
