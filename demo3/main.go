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
	url := launcher.New().Headless(false).
		Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0").
		Set("disable-blink-features", "AutomationControlled").MustLaunch()
	//建立连接
	browser := rod.New().ControlURL(url).MustConnect()
	//延迟关闭
	defer browser.Close()
	listPage := browser.MustPage("https://sale.1688.com/factory/u0vjcc4j.html?spm=a260k.home2025.centralDoor.ddoor.66333597BBbHgE&topOfferIds=1005591171200")
	listPage.Timeout(20 * time.Second)
	listPage.MustWaitLoad()

	err := listPage.MustElement(".offerItem").WaitVisible()
	if err != nil {
		fmt.Println("加载商品失败")
		return
	}
	listPage.Timeout(10 * time.Second)
	listPage.Mouse.Scroll(0, 500, 0)
	elements := listPage.MustElements(".offerItem")

	fmt.Printf("抓取到商品数量%d\n", len(elements))
	if len(elements) <= 0 {
		fmt.Println("未抓取到商品")
		return
	}
	time.Sleep(5 * time.Second)
	for i, item := range elements {
		if !item.MustVisible() {
			continue
		}
		priceSpans := item.MustElements("span.text")
		var priceParts []string
		for _, span := range priceSpans {
			text := span.MustText()
			// 判断是否是价格部分（包含数字或¥符号）
			if strings.Contains(text, "¥") || strings.Contains(text, ".") ||
				(len(text) > 0 && text[0] >= '0' && text[0] <= '9') {
				priceParts = append(priceParts, text)
			}
		}
		price := strings.Join(priceParts, "")

		// 如果上面没拿到，尝试直接取父容器的全部文本
		if price == "" {
			if priceBox, err := item.Element("div[class*='price'], div[class*='Price']"); err == nil {
				price, _ = priceBox.Text()
			}
		}

		// 获取标题
		title := ""
		if titleEl, err := item.Element(".offerTitle, .title"); err == nil {
			title, _ = titleEl.Text()
		}
		src := ""
		if srcEl, err := item.Element(".offerImg"); err == nil {
			attribute, err := srcEl.Attribute("src")
			if err == nil && attribute != nil {
				src = fmt.Sprintf("https:%s", *attribute)
			}
		}
		fmt.Printf("[%d] 标题: %s | 价格: %s | 地址: %s\n", i+1, title, price, src)

	}
}
