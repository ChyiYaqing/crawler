package main

import (
	"fmt"
	"time"

	"github.com/chyiyaqing/crawler/collect"
	"github.com/chyiyaqing/crawler/proxy"
)

// [\s\S] 指代的是任意字符串
// * 代表将前面任意字符匹配0次或者无数次
// ? 代表非贪婪匹配
//var headerRe = regexp.MustCompile(`<div class="small_cardcontent__BTALp"[\s\S]*?<h2>([\s\S]*?)</h2>`)

func main() {

	proxyURLs := []string{"http://127.0.0.1:8888", "http://127.0.0.1:8889"}
	p, err := proxy.RoundRobinProxySwitcher(proxyURLs...)
	if err != nil {
		fmt.Println("RoundRobinProxySwitcher failed")
	}

	// 澎湃信息
	//url := "https://www.thepaper.cn/"
	// 豆瓣读书
	// url := "https://book.douban.com/subject/1060068/"
	url := "https://google.com"
	var f collect.Fetcher = collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
		Proxy:   p,
	}

	body, err := f.Get(url)
	if err != nil {
		fmt.Printf("read content failed:%v\n", err)
		return
	}
	fmt.Println(string(body))
}
