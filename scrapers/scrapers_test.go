package scrapers_test

import (
	"scraper/scrapers"
	"testing"
)

func TestGetUrls(t *testing.T) {
	t.Parallel()
	tt := []struct {
		testName    string
		companyName string
		want        string
	}{{
		testName:    "Url without www",
		companyName: "kfc ",
		want:        "https://jfood.kfc.ug/",
	},
		{
			testName:    "Url with .org",
			companyName: "innovex",
			want:        "https://innovex.org/",
		},
		{
			testName:    "Url with www",
			companyName: "netflix",
			want:        "https://www.netflix.com/",
		},
		{
			testName:    "Company name in caps",
			companyName: "NETFLIX",
			want:        "https://www.netflix.com/",
		},
	}
	for _, testCase := range tt {
		testCase := testCase
		t.Run(testCase.testName, func(t *testing.T) {
			t.Parallel()
			got := scrapers.GetUrls(testCase.companyName)
			if got != testCase.want {
				t.Errorf("Got URL:%s, Expected URL: %s", got, testCase.want)
			}
		})

	}
}

func TestAboutUs(t *testing.T) {
	t.Parallel()

	tt := []struct {
		testName   string
		companyURL string
		want       string
	}{
		{
			"With about-us",
			"https://innovex.org/",
			"https://innovex.org/about-us/",
		},
		{
			"With contactus",
			"https://www.netflix.com/ug/",
			"https://help.netflix.com/contactus",
		},
		{
			"With about",
			"http://www.netlabsug.org/",
			"http://www.netlabsug.org/website/about",
		},
	}

	for _, testCase := range tt {
		testCase := testCase
		t.Run(testCase.testName, func(t *testing.T) {
			t.Parallel()

			got := scrapers.AboutUs(testCase.companyURL)
			if got != testCase.want {
				t.Errorf("Got URL:%s, Expected URL: %s", got, testCase.want)
			}

		})
	}

}
