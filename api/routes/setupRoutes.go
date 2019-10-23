package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andytan0727/newster/api/news"
	"github.com/gorilla/mux"
)

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
	"dev-to": {
		Name: "DEV TO",
		URL:  "https://dev.to/",
	},
	"github-trending": {
		Name: "Github Trending",
		URL:  "https://github.com/trending",
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

	scraper = news.CSDNScraper{Request: news.RequestNews{}, URL: csdn.URL}

	if csdnNews, err = scraper.Scrap(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if csdnNewsJSON, err = json.Marshal(news.JSONResp{
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

	scraper = news.CSSTricksScraper{Request: news.RequestNews{}, URL: cssTricks.URL}

	if cssTricksArticles, err = scraper.Scrap(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if cssTricksJSON, err = json.Marshal(news.JSONResp{
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

// devToScrapRouteHandler is a handler to get
// latest articles (news) data from dev.to
func devToScrapRouteHandler(w http.ResponseWriter, r *http.Request) {
	SetCORSHeaders(w)

	var (
		scraper       news.Scraper
		devToArticles []news.News
		devToJSON     []byte
		err           error
	)
	devTo := sites["dev-to"]

	log.Printf("Scraping %s...", devTo.Name)

	scraper = news.DevToScraper{Request: news.RequestNews{}, URL: devTo.URL}

	if devToArticles, err = scraper.Scrap(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if devToJSON, err = json.Marshal(news.JSONResp{
		News: devToArticles,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(devToJSON)

	// log success message if scrap ends with no err
	if err == nil {
		log.Printf("Successfully scraped %s", devTo.Name)
	}
}

// githubTrendingScrapRouteHandler is a handler to get
// latest articles (news) data from github trending
func githubTrendingScrapRouteHandler(w http.ResponseWriter, r *http.Request) {
	SetCORSHeaders(w)

	var (
		scraper                 news.Scraper
		githubTrendingRepos     []news.News
		githubTrendingReposJSON []byte
		err                     error
	)
	githubTrending := sites["github-trending"]

	log.Printf("Scraping %s...", githubTrending.Name)

	scraper = news.GithubTrendingScraper{Request: news.RequestNews{}, URL: githubTrending.URL}

	if githubTrendingRepos, err = scraper.Scrap(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if githubTrendingReposJSON, err = json.Marshal(news.JSONResp{
		News: githubTrendingRepos,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(githubTrendingReposJSON)

	// log success message if scrap ends with no err
	if err == nil {
		log.Printf("Successfully scraped %s", githubTrending.Name)
	}
}

// SetupRoutes setup routes for all routes in this api
func SetupRoutes(r *mux.Router) {
	apiRoute := r.PathPrefix("/api").Subrouter()
	apiRoute.Methods(http.MethodGet).Path("/csdn").HandlerFunc(csdnScrapRouteHandler)
	apiRoute.Methods(http.MethodGet).Path("/css-tricks").HandlerFunc(cssTricksScrapRouteHandler)
	apiRoute.Methods(http.MethodGet).Path("/dev-to").HandlerFunc(devToScrapRouteHandler)
	apiRoute.Methods(http.MethodGet).Path("/gh-trending").HandlerFunc(githubTrendingScrapRouteHandler)
}
