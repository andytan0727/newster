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
func (r Request) RequestData() (io.ReadCloser, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	client := &http.Client{
		Timeout: timeout,
	}

	if req, err = http.NewRequest(http.MethodGet, r.URL, nil); err != nil {
		return nil, ScrapError(r.URL, err)
	}

	AddRequestHeaders(req)

	if resp, err = client.Do(req); err != nil {
		return nil, ScrapError(r.URL, err)
	}

	return resp.Body, nil
}

// Scrap scraps data from CSDN
func (s CSDNScraper) Scrap() ([]News, error) {
	var (
		body io.ReadCloser
		err  error
		news []News
	)

	if body, err = s.RequestData(); err != nil {
		return []News{}, ScrapError(s.URL, err)
	}

	defer body.Close()

	document, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		return []News{}, QueryDocumentError(s.URL, err)
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

// Scrap scraps data from CSS-Trick
func (s CSSTricksScraper) Scrap() ([]News, error) {
	var (
		body io.ReadCloser
		err  error
		news []News
	)

	if body, err = s.RequestData(); err != nil {
		return []News{}, ScrapError(s.URL, err)
	}

	defer body.Close()

	document, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		return []News{}, QueryDocumentError(s.URL, err)
	}

	document.Find(".article-article").Each(func(i int, sel *goquery.Selection) {
		a := sel.Find("h2 a").First()
		url, exist := a.Attr("href")
		newsTitle := strings.TrimSpace(a.Text())

		if len(newsTitle) != 0 && exist {
			news = append(news, News{Title: newsTitle, URL: url})
		}
	})

	return news, nil
}
