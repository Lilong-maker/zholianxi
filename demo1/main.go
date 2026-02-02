package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	word := launcher.New().
		Headless(true).
		Set("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36").
		MustLaunch()
	ff := rod.New().ControlURL(word).MustConnect()
	defer ff.MustClose()

	page := ff.MustPage("https://sale.1688.com/factory/u0vjcc4j.html?spm=a260k.home2025.centralDoor.ddoor.66333597BBbHgE&topOfferIds=1005591171200")
	page.MustWaitLoad()
	err := page.Timeout(15 * time.Second).MustElement(`.offerTitle`).WaitVisible()
	if err != nil {
		return
	}

	time.Sleep(2 * time.Second)
	num := page.MustElements(`.offerTitle`)
	for i, e := range num {
		if e.MustVisible() {
			text := e.MustText()
			fmt.Printf("%d. %s\n", i+1, text)
		}
	}
}
