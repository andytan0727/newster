package news

// News defines struct holding information scrapped: title and url of a news
type News struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Request consists of a URL of website to request,
// and a RequestData method
type Request struct {
	URL string
}

// CSDNScraper contains method to scrap CSDN and an
// extended Request struct to perform HTTP request
type CSDNScraper struct {
	Request
}

// CSSTricksScraper contains method to scrap CSS-Trick
// and an extended Request struct to perform HTTP request
type CSSTricksScraper struct {
	Request
}
