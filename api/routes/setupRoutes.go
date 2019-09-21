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

// csdnScrapRouteHandler is a handler to get latest
// news data from CSDN
func csdnScrapRouteHandler(w http.ResponseWriter, r *http.Request) {
	SetCORSHeaders(w)
	log.Println("Scraping CSDN...")

	var (
		scraper      news.Scraper
		csdnNews     []news.News
		csdnNewsJSON []byte
		err          error
	)

	scraper = news.CSDNScraper{Request: news.Request{URL: "https://www.csdn.net/"}}

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
		log.Println("Successfully scraped CSDN...")
	}
}

// SetupRoutes setup routes for all routes in this api
func SetupRoutes(r *mux.Router) {
	apiRoute := r.PathPrefix("/api")
	apiRoute.Methods(http.MethodGet).Path("/csdn").HandlerFunc(csdnScrapRouteHandler)
}
