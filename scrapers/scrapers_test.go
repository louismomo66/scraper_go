package scrapers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"scraper/scrapers"
	"testing"
)
func TestGetUrls(t *testing.T) {
	tt := []struct {
		company string
		htmlResponse string
		expectedURL string
	}{{
		"codebits",`<html>
		<body>
		<a href="/url?q=http://codebits.com">Example</a>
		<a href="https://codebits.com">Example</a>
		</body>
	</html>`,"https://codebits.in/",
	},
}

for _,tc := range tt {
	t.Run(tc.company,func(t *testing.T){
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			io.WriteString(w,tc.htmlResponse)
		}))
		defer server.Close()
		// baseURL := server.URL
		result := scrapers.GetUrls(tc.company)
		if result != tc.expectedURL {
			t.Errorf("For %s: Expected Url: %s, but got: %s", tc.company,tc.expectedURL,result)
		}
	})
}

}