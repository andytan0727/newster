package news

// News defines struct holding information scrapped: title and url of a news
type News struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// JSONResp is JSON response containing news
type JSONResp struct {
	News []News `json:"news"`
}

// RequestNews implements Request interface that consists
// a RequestData method to be used in Scrapers
type RequestNews struct{}

// CSDNScraper contains method to scrap CSDN and an
// extended Request interface to perform HTTP request
type CSDNScraper struct {
	URL string
	Request
}

// CSSTricksScraper contains method to scrap CSS-Trick
// and an extended Request interface to perform HTTP request
type CSSTricksScraper struct {
	URL string
	Request
}

// DevToScraper contains method to scrap dev.to
// and an extended Request interface to perform HTTP request
type DevToScraper struct {
	URL string
	Request
}
