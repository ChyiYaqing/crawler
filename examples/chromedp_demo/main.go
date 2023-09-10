package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// 1. 创建谷歌浏览器实例
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 2. 设置context超时时间
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// 3. 爬取页面,等待某个元素出现，接着模拟鼠标点击，最后获取数据
	var example string
	// chromedp.Run 执行多个action
	// action指的是爬取、等待、点击、获取数据这样的行为
	// task指的是一个任务
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://pkg.go.dev/time`),           // 爬取指定的网址
		chromedp.WaitVisible(`body > footer`),                  // 等待当前标签可见，body > footer标签可见，代表正文已经加载完毕
		chromedp.Click(`#example-After`, chromedp.NodeVisible), // 模拟对某一个标签的点击事件
		chromedp.Value(`#example-After textarea`, &example),    // 用于获取指定标签的数据
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}
