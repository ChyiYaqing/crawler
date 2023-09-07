package main

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/chyiyaqing/crawler/collect"
)

// [\s\S] 指代的是任意字符串
// * 代表将前面任意字符匹配0次或者无数次
// ? 代表非贪婪匹配
//var headerRe = regexp.MustCompile(`<div class="small_cardcontent__BTALp"[\s\S]*?<h2>([\s\S]*?)</h2>`)

func main() {
	url := "https://www.thepaper.cn/"
	var f collect.Fetcher = collect.BaseFetch{}
	body, err := f.Get(url)

	if err != nil {
		fmt.Printf("read content failed:%v\n", err)
		return
	}

	// 加载HTML文档
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("read content failed:%v\n", err)
	}

	doc.Find("div.small_cardcontent__BTALp a[target=_blank] h2").Each(func(i int, s *goquery.Selection) {
		// 获取匹配标签中的文本
		title := s.Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})

}
