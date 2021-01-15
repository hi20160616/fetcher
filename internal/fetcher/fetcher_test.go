package fetcher

import (
	"log"
	"testing"
	"time"
)

func TestCrawl(t *testing.T) {
	for {
		BreadthFirst(Crawl, []string{
			// "https://www.dwnews.com",
			// "https://www.zaobao.com/realtime/world",
			// "https://www.zaobao.com/news/world",
			// "https://www.voachinese.com",
			// "https://news.ltn.com.tw/list/breakingnews",
			// "https://www.cna.com.tw/list/aall.aspx",
			"https://www.bbc.com/zhongwen/simp/topics/ck2l9z0em07t",
		})

		log.Println("Sleep a sec ...")
		time.Sleep(5 * time.Minute) // only useful by goroutine
	}
}
