package news

// Scraper defines a standard scrap function to all scrapers
type Scraper interface {
	Scrap() ([]News, error)
}
