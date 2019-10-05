package news

import "io"

// Scraper defines a standard scrap function to all scrapers
type Scraper interface {
	Scrap() ([]News, error)
}

// Request defines a standard HTTP Request interface
// for Scrapers that contains RequestData function
type Request interface {
	RequestData(string) (io.ReadCloser, error)
}
