package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/giskook/shoe/conf"
)

func NoHeadless(a *chromedp.ExecAllocator) {
	chromedp.Flag("headless", false)(a)
	// Like in Puppeteer.
	chromedp.Flag("hide-scrollbars", false)(a)
	chromedp.Flag("mute-audio", false)(a)
	chromedp.Flag("disable-web-security", true)(a)
}
func main() {
	if time.Now().Unix() > 1569686400 {
		return
	}
	conf := conf.Parse()
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		NoHeadless,
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// ensure that the browser process is started
	err := chromedp.Run(taskCtx,
		//	chromedp.Navigate(`https://detail.tmall.com/item.htm?spm=a1z10.4-b-s.w5003-21968665072.1.21abac7f24HhYO&id=562079049525&scene=taobao_sho`),
		chromedp.Navigate(conf.Url),
		// wait for footer element is visible (ie, page is loaded)
	)
	log.Println("login,select your size then type any key... ...")
	var flag string
	fmt.Scanln(&flag)
	if err != nil {
		panic(err)
	}
	err = chromedp.Run(taskCtx,
		chromedp.Click(conf.ButtonID),
	)
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
}
