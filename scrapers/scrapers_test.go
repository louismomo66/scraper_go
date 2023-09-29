package scrapers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUrls(t *testing.T) {
	testTable := []struct {
		name          string
		companyName   string
		serverHandler http.HandlerFunc
		expectedUrl   string
		expectedErr   bool
	}{
		{
			name:        "Successful",
			companyName: "netlabs",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `<a href="/url?q=http://www.netlabsug.org/">Netlabs UG</a>`)
			},
			expectedUrl: "http://www.netlabsug.org/",
			expectedErr: false,
		},

		{
			name:        "HTTP error",
			companyName: "dummy",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "test error", http.StatusInternalServerError)
			},
			expectedUrl: "",
			expectedErr: true,
		},
		{
			name:        "Non- successful status code",
			companyName: "any",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
			expectedUrl: "",
			expectedErr: true,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			ts := httptest.NewServer(tc.serverHandler)
			defer ts.Close()
			url, err := GetUrls(ts.URL, tc.companyName)
			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedUrl, url)
			}
		})
	}

}

func TestAboutUs(t *testing.T) {
	testTable := []struct {
		name            string
		serverHandler   http.HandlerFunc
		expectedAboutUs string
		expectedErr     error
	}{
		{
			name: "Extract_aboutUs_Link",
			serverHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `<a href="/about-us">About Us</a>`)
			}),
		},
		{
			name: "HTTP_error",
			serverHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}),
			expectedErr: fmt.Errorf("unexpected status code: 500"),
		},
		{
			name: "Non-successful_status_code",
			serverHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Not Found", http.StatusNotFound)
			}),
			expectedErr: fmt.Errorf("unexpected status code: 404"),
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			ts := httptest.NewServer(tc.serverHandler)
			defer ts.Close()
			if tc.name == "Extract_aboutUs_Link" {
				tc.expectedAboutUs = ts.URL + "/about-us"
			} else if tc.name == "No_relevant_link" {
				tc.expectedAboutUs = ""
			}
			got, err := AboutUs(ts.URL)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Fatalf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
			if got != tc.expectedAboutUs {
				t.Fatalf("Expected AboutUS URL: %s, got: %s", tc.expectedAboutUs, got)
			}
		})
	}
}

func TestExtractEmail(t *testing.T) {
	tt := []struct {
		name          string
		content       string
		serverHandler http.HandlerFunc
		expectedEmail string
		expectedErr   bool
	}{
		{
			name:    "Success",
			content: "dummy",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `Contact us at: contact@dummy.com`)
			},
			expectedEmail: "contact@dummy.com",
			expectedErr:   false,
		},
		{
			name:    "HTTP error on fetching content",
			content: "dummy",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "test error", http.StatusInternalServerError)
			},
			expectedEmail: "",
			expectedErr:   true,
		},
		{
			name:    "Non-successful status code",
			content: "dummy",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
			expectedEmail: "",
			expectedErr:   true,
		},
		{
			name:    "No email in content",
			content: "dummy",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `Hello World`)
			},
			expectedEmail: "",
			expectedErr:   false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ts := httptest.NewServer(tc.serverHandler)
			defer ts.Close()

			email, err := ExtractEmail(ts.URL)
			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedEmail, email)
			}
		})
	}
}
