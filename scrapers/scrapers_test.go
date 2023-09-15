package scrapers_test

import (
	"fmt"
	"scraper/scrapers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUrls(t *testing.T) {
	t.Parallel()
	type args struct {
		companyName string
	}
	tt := []struct {
		testName  string
		nameAgrs  args
		want      string
		wantedErr error
	}{
		{
			"Url without www",
			args{"mukwano"},
			"https://www.mukwano.com/",
			nil,
		},
		// {
		// 	"Url with .org",
		// 	"innovex",
		// 	"https://www.innovex-inc.com/",
		// },
		{
			"Url with www",
			args{"netflix"},
			"https://www.netflix.com/",
			nil,
		},
		{
			"Company name in caps",
			args{"NETFLIX"},
			"https://www.netflix.com/",
			nil,
		},
		{
			"Company name in two words",
			args{"roke telecom"},
			"https://www.roketelkom.co.ug/",
			nil,
		},
	}
	for i := range tt { //nolint
		i := i
		t.Run(tt[i].testName, func(t *testing.T) {
			t.Parallel()
			got, err := scrapers.GetUrls(tt[i].nameAgrs.companyName)
			if err != nil && tt[i].wantedErr == nil {
				// t.Errorf("Got URL:%s, Expected URL: %s", got, tt[i].want)
				assert.Fail(t, fmt.Sprintf("Error not expected but got one:\n"+"error: %q", err))
				return
			}
			if tt[i].wantedErr != nil {
				assert.EqualError(t, err, tt[i].wantedErr.Error())
				return
			}
			assert.Equal(t, tt[i].want, got)
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

	for i := range tt {
		i := i
		t.Run(tt[i].testName, func(t *testing.T) {
			t.Parallel()

			got, err := scrapers.AboutUs(tt[i].companyURL)
			if err != nil && tt[i].wantedErr == nil {
				// t.Errorf("Got URL:%s, Expected URL: %s", got, testCase.want)
				assert.Fail(t, fmt.Sprintf("Error not expected but got one:\n"+"error: %q", err))
				return
			}
			if tt[i].wantedErr != nil {
				assert.EqualError(t, err, tt[i].wantedErr.Error())
				return
			}
			assert.Equal(t, tt[i].want, got)

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
