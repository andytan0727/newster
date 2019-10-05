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
func (r RequestNews) RequestData(url string) (io.ReadCloser, error) {
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

// Scrap scraps data from CSDN
func (s CSDNScraper) Scrap() ([]News, error) {
	var (
		body io.ReadCloser
		err  error
		news []News
	)

	if body, err = s.RequestData(s.URL); err != nil {
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

	if body, err = s.RequestData(s.URL); err != nil {
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

// Scrap from DevToScraper scraps data from dev.to
func (s DevToScraper) Scrap() ([]News, error) {
	var (
		body io.ReadCloser
		err  error
		news []News
	)

	if body, err = s.RequestData(s.URL); err != nil {
		return []News{}, ScrapError(s.URL, err)
	}

	defer body.Close()

	document, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		return []News{}, QueryDocumentError(s.URL, err)
	}

	document.Find(".index-article-link").Each(func(i int, sel *goquery.Selection) {
		a := sel
		url, exist := a.Attr("href")
		titleContent := sel.Find("h3")

		// remove span that contains hashtags
		titleContent.Find("span").Remove()

		newsTitle := strings.TrimSpace(titleContent.Text())

		if len(newsTitle) != 0 && exist {
			fullURL := "https://dev.to" + url
			news = append(news, News{Title: newsTitle, URL: fullURL})
		}
	})

	return news, nil
}
