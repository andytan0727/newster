package news

import (
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"
	timeout   = time.Duration(5 * time.Second)
)

var client = &http.Client{
	Timeout: timeout,
}

// ScrapCSDN scraps data from CSDN
func ScrapCSDN(url string) ([]News, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return ShowScrapError(url, err)
	}

	AddRequestHeaders(req)
	resp, err := client.Do(req)

	if err != nil {
		return ShowScrapError(url, err)
	}

	defer resp.Body.Close()
	var news []News
	document, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return ShowScrapError(url, err)
	}

	document.Find(".title").Each(func(i int, sel *goquery.Selection) {
		a := sel.Find("h2 a").First()
		url, exist := a.Attr("href")
		newsTitle := strings.TrimSpace(a.Text())

		if len(newsTitle) != 0 && exist {
			news = append(news, News{Title: newsTitle, URL: url})
		}
	})

	return news, nil
}
