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
		wantedErr   error
	}{
		// {
		// 	"Url without www",
		// 	"kfc ",
		// 	"https://jfood.kfc.ug/",
		// },
		// {
		// 	"Url with .org",
		// 	"innovex",
		// 	"https://www.innovex-inc.com/",
		// },
		{
			"Url with www",
			"netflix",
			"https://www.netflix.com/",
			nil,
		},
		{
			"Company name in caps",
			"NETFLIX",
			"https://www.netflix.com/",
			nil,
		},
		{
			"Company name in two words",
			"roke telecom",
			"https://www.roketelkom.co.ug/",
			nil,
		},
	}
	for _, testCase := range tt {
		testCase := testCase
		t.Run(testCase.testName, func(t *testing.T) {
			t.Parallel()
			got, err := scrapers.GetUrls(testCase.companyName)
			if err != nil && got != testCase.want {
				t.Errorf("Got URL:%s, Expected URL: %s", got, testCase.want)
				return
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
		wantedErr  error
	}{
		{
			"With about-us",
			"https://innovex.org/",
			"https://innovex.org/about-us/",
			nil,
		},
		// {
		// 	"With contactus",
		// 	"https://www.netflix.com/",
		// 	"https://help.netflix.com/contactus",
		//  nil,
		// },
		{
			"With about",
			"http://www.netlabsug.org/",
			"http://www.netlabsug.org/website/about",
			nil,
		},
	}

	for _, testCase := range tt {
		testCase := testCase
		t.Run(testCase.testName, func(t *testing.T) {
			t.Parallel()

			got, err := scrapers.AboutUs(testCase.companyURL)
			if err != nil && got != testCase.want {
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
		wantedErr  error
	}{
		{
			"Extraction from contact-us",
			"https://ug.liquidhome.tech/contact-us",
			"ugsales@liquidtelecom.com",
			nil,
		},
		{
			"no email found",
			"https://innovex.org/about-us/",
			"",
			nil,
		},
	}

	for _, testCase := range tt {
		testCase := testCase
		t.Run(testCase.testName, func(t *testing.T) {
			t.Parallel()

			got, err := scrapers.ExtractEmail(testCase.aboutUsURL)
			if err != nil && got != testCase.want {
				t.Errorf("Got %s, Expected %s", got, testCase.want)
			}
		})
	}
}
