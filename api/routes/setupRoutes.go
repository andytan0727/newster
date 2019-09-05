package routes

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("Scraping CSDN...")
	csdnNews, err := news.ScrapCSDN("https://www.csdn.net/")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	csdnNewsJSON, err := json.Marshal(NewsJSONResp{
		News: csdnNews,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(csdnNewsJSON)
}

// SetupRoutes setup routes for all routes in this api
func SetupRoutes(r *mux.Router) {
	apiRoute := r.PathPrefix("/api")
	apiRoute.Methods(http.MethodGet).Path("/csdn").HandlerFunc(csdnScrapRouteHandler)
}
