package news

import (
	"fmt"
	"log"
	"net/http"
)

// AddRequestHeaders add some headers to a request,
// particularly User-Agent and Upgrade-Insecure-Requests headers
func AddRequestHeaders(req *http.Request) {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Upgrade-Insecure-Requests", `1`)
}

// ScrapError prints and return error encountered when scraping
func ScrapError(url string, err error) error {
	msg := fmt.Sprintf("Error scraping from %s: %v", url, err.Error())
	log.Println(msg)
	return fmt.Errorf(msg)
}

// QueryDocumentError prints and returns error encountered when querying
// html document with goquery
func QueryDocumentError(url string, err error) error {
	msg := fmt.Sprintf("Error query document from %s: %v", url, err.Error())
	log.Println(msg)
	return fmt.Errorf(msg)
}
