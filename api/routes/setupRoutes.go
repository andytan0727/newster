package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andytan0727/newster/api/news"
	"github.com/gorilla/mux"
)

// NewsJSONResp is JSON response containing news
type NewsJSONResp struct {
	News []news.News `json:"news"`
}

type siteInfo struct {
	Name string
	URL  string
}

// sites stored all information of the sites to be scraped
var sites = map[string]siteInfo{
	"csdn": {
		Name: "CSDN",
		URL:  "https://www.csdn.net/",
	},
	"css-tricks": {
		Name: "CSS-Tricks",
		URL:  "https://css-tricks.com/archives/",
	},
}

// csdnScrapRouteHandler is a handler to get latest
// news data from CSDN
func csdnScrapRouteHandler(w http.ResponseWriter, r *http.Request) {
	SetCORSHeaders(w)

	var (
		scraper      news.Scraper
		csdnNews     []news.News
		csdnNewsJSON []byte
		err          error
	)
	csdn := sites["csdn"]

	log.Printf("Scraping %s...", csdn.Name)

	scraper = news.CSDNScraper{Request: news.Request{URL: csdn.URL}}

	if csdnNews, err = scraper.Scrap(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if csdnNewsJSON, err = json.Marshal(NewsJSONResp{
		News: csdnNews,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(csdnNewsJSON)

	// log success message if scrap ends with no err
	if err == nil {
		log.Printf("Successfully scraped %s...", csdn.Name)
	}
}

// cssTricksScrapRouteHandler is a handler to get
// latest articles (news) data from CSS-Trick
func cssTricksScrapRouteHandler(w http.ResponseWriter, r *http.Request) {
	SetCORSHeaders(w)

	var (
		scraper           news.Scraper
		cssTricksArticles []news.News
		cssTricksJSON     []byte
		err               error
	)
	cssTricks := sites["css-tricks"]

	log.Printf("Scraping %s...", cssTricks.Name)

	scraper = news.CSSTricksScraper{Request: news.Request{URL: cssTricks.URL}}

	if cssTricksArticles, err = scraper.Scrap(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if cssTricksJSON, err = json.Marshal(NewsJSONResp{
		News: cssTricksArticles,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(cssTricksJSON)

	// log success message if scrap ends with no err
	if err == nil {
		log.Printf("Successfully scraped %s", cssTricks.Name)
	}
}

// SetupRoutes setup routes for all routes in this api
func SetupRoutes(r *mux.Router) {
	apiRoute := r.PathPrefix("/api").Subrouter()
	apiRoute.Methods(http.MethodGet).Path("/csdn").HandlerFunc(csdnScrapRouteHandler)
	apiRoute.Methods(http.MethodGet).Path("/css-tricks").HandlerFunc(cssTricksScrapRouteHandler)
}
