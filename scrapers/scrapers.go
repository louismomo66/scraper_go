package scrapers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetUrls(companyName string) (string, error) {
	escapedCompanyName := strings.ReplaceAll(companyName, " ", "+")
	pageLink := fmt.Sprintf("http://google.com/search?q=%s", escapedCompanyName)
	resp, httpErr := http.Get(pageLink) //nolint
	if httpErr != nil {                 //nolint
		err := fmt.Errorf("an error occurred trying to scrapper google for %s %w", companyName, httpErr)
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err //nolint
	}
	foundURL := ""
	doc.Find("body a").Each(func(index int, item *goquery.Selection) { //nolint
		if foundURL == "" {
			linkTag := item
			link, _ := linkTag.Attr("href")

			if strings.HasPrefix(link, "/url?q=") {
				link = strings.TrimPrefix(link, "/url?q=")
				link = strings.Split(link, "&")[0]
				foundURL = link
			}
		}
	})

	return foundURL, nil
}

func ExtractEmail(content string) (string, error) {
	resp, err := http.Get(content) //nolint
	if err != nil {
		log.Println(err)
		return "", err //nolint
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil { //nolint
		log.Println(err)
		return "", err //nolint
	}

	// re := regexp.MustCompile(`[\w\.-]+@[\w\.-]+`)
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	match := re.FindString(string(data))
	return match, nil //nolint
}

func AboutUs(companyURL string) (string, error) {
	aboutUsURL := ""
	resp, err := http.Get(companyURL) //nolint
	if err != nil {                   //nolint
		log.Printf("Error fetching %s: %v\n", companyURL, err)
		return "", err //nolint
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading respinse body: %v\n", err)
		return "", err //nolint
	}
	re := regexp.MustCompile(`(?i)<a[^>]+href=["']([^"']+)["'][^>]*>(?:\s*about\s*us\s*|about|contact\s*us)\s*</a>`)
	match := re.FindStringSubmatch(string(data))
	if len(match) > 1 { //nolint
		aboutUsURL = match[1]
	}
	if !strings.HasPrefix(aboutUsURL, "http") { //nolint
		aboutUsURL = companyURL + aboutUsURL
	}
	return aboutUsURL, nil //nolint
}
