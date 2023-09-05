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
	}{
		{
			"Url without www",
			"kfc ",
			"https://jfood.kfc.ug/",
		},
		// {
		// 	"Url with .org",
		// 	"innovex",
		// 	"https://www.innovex-inc.com/",
		// },
		{
			"Url with www",
			"netflix",
			"https://www.netflix.com/",
		},
		{
			"Company name in caps",
			"NETFLIX",
			"https://www.netflix.com/",
		},
		{
			"Company name in two words",
			"roke telecom",
			"https://www.roketelkom.co.ug/",
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
		// {
		// 	"With contactus",
		// 	"https://www.netflix.com/",
		// 	"https://help.netflix.com/contactus",
		// },
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

func TestExtructEmail(t *testing.T) {
	t.Parallel()

	tt := []struct {
		testName   string
		aboutUsURL string
		want       string
	}{
		{
			"Extraction from contact-us",
			"https://ug.liquidhome.tech/contact-us",
			"ugsales@liquidtelecom.com",
		},
		{
			"no email found",
			"https://innovex.org/about-us/",
			"",
		},
	}

	for _, testCase := range tt {
		testCase := testCase
		t.Run(testCase.testName, func(t *testing.T) {
			t.Parallel()

			got := scrapers.ExtractEmail(testCase.aboutUsURL)
			if got != testCase.want {
				t.Errorf("Got %s, Expected %s", got, testCase.want)
			}
		})
	}
}
