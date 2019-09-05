package news

import (
	"fmt"
	"net/http"
)

// AddRequestHeaders add some headers to a request,
// particularly User-Agent and Upgrade-Insecure-Requests headers
func AddRequestHeaders(req *http.Request) {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Upgrade-Insecure-Requests", `1`)
}

// ShowScrapError shows error encountered when scraping
func ShowScrapError(url string, err error) ([]News, error) {
	fmt.Printf("Error scraping from %s: %v\n", url, err)
	return []News{}, err
}
