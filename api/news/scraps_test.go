package news

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
	"testing"
)

type mockedRequest struct{}

// RequestData is a function mocking functionality of
// RequestNews's RequestData with opening mocked HTML
// from file system
func (m mockedRequest) RequestData(url string) (io.ReadCloser, error) {
	f, err := os.Open(url)

	if err != nil {
		return nil, err
	}

	return f, nil
}

var testCases = []struct {
	scraper        Scraper
	wantNewsLength int
}{
	{scraper: CSDNScraper{Request: mockedRequest{}, URL: "../mocks/mockCSDNNews.html"}, wantNewsLength: 29},

	{scraper: CSSTricksScraper{Request: mockedRequest{}, URL: "../mocks/mockCssTricksNews.html"}, wantNewsLength: 10},

	{scraper: DevToScraper{Request: mockedRequest{}, URL: "../mocks/mockDevToNews.html"}, wantNewsLength: 35},
}

func TestScrapersScrap(t *testing.T) {
	var (
		news        []News
		scraperType reflect.Type
		scraperName string
		err         error
	)

	for _, tc := range testCases {
		scraperVal := reflect.ValueOf(tc.scraper)

		// if its a pointer, resolve its value
		if scraperVal.Kind() == reflect.Ptr {
			scraperVal = reflect.Indirect(scraperVal)
		}

		scraperType = scraperVal.Type()
		scraperName = scraperType.Name()

		t.Logf("Testing scraper %s...", scraperName)

		t.Run("Test_URL_field_exists", func(t *testing.T) {
			_, got := scraperType.FieldByName("URL")
			want := true

			if got != want {
				t.Errorf("URL does not exists: got = %v, want = %v", got, want)
			}
		})

		t.Run("Test_News_Length", func(t *testing.T) {
			if news, err = tc.scraper.Scrap(); err != nil {
				t.Errorf("Error in Scrap: %v", err)
			}

			gotNewsLength := len(news)

			if gotNewsLength != tc.wantNewsLength {
				t.Errorf("news length mismatch: want: %d, got: %d", tc.wantNewsLength, gotNewsLength)
			}

		})

		t.Run("Test_JSON_Marshalling", func(t *testing.T) {
			if _, err = json.Marshal(JSONResp{
				News: news,
			}); err != nil {
				t.Errorf("Error when marshalling to JSON: %v", err)
			}
		})
	}
}
