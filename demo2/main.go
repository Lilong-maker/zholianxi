package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	yy := launcher.New().
		Headless(true).
		Set("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36").
		MustLaunch()
	uu := rod.New().ControlURL(yy).MustConnect()

	page := uu.MustPage("https://sale.1688.com/factory/u0vjcc4j.html?spm=a260k.home2025.centralDoor.ddoor.66333597BBbHgE&topOfferIds=1005591171200")

	err := page.Timeout(30 * time.Second).MustElement(".offerImg").WaitVisible()
	if err != nil {
		return
	}

	time.Sleep(2 * time.Second)
	pp := page.MustElements(".offerImg")
	for i, e := range pp {
		if e.MustVisible() {
			attribute, err := e.Attribute("src")
			if err == nil && attribute != nil {
				fmt.Printf("%d. %s\n", i+1, *attribute)
			}
		}
	}
}
