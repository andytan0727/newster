package news

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const timeout = time.Duration(10 * time.Second)

// RequestData sends http request to particular site
// and return a response body
func RequestData(url string) (io.ReadCloser, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	client := &http.Client{
		Timeout: timeout,
	}

	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, ScrapError(url, err)
	}

	AddRequestHeaders(req)

	if resp, err = client.Do(req); err != nil {
		return nil, ScrapError(url, err)
	}

	return resp.Body, nil
}

// ScrapCSDN scraps data from CSDN
func ScrapCSDN(url string, body io.ReadCloser) ([]News, error) {
	defer body.Close()

	var news []News
	document, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		return []News{}, QueryDocumentError(url, err)
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
