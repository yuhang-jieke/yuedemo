package getres

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func GetImg(address string, selector string) {
	url := launcher.New().
		Headless(false).
		Set("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0").
		MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()
	page := browser.MustPage(address)
	page.MustWaitLoad()
	err := page.Timeout(2 * time.Minute).MustElement(selector).WaitVisible()
	if err != nil {
		fmt.Println("等待元素超时")
		return
	}
	time.Sleep(5 * time.Second)
	element := page.MustElements(selector)
	fmt.Printf("获取到%d个元素\n", len(element))
	for i, e := range element {
		if e.MustVisible() {
			src, err := e.Attribute("src")
			if err == nil && src != nil {
				fmt.Printf("%d.%s\n", i+1, *src)
			}

		}
	}
}
