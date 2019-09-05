package news

import "time"

// News defines struct holding information scrapped: title and url of a news
type News struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Scraper contains information necessary for a req
// to be success
type Scraper struct {
	UserAgent string
	Timeout   time.Duration
}
