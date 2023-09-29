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

func GetUrls(baseURL, companyName string) (string, error) {
	fullCompanyName := strings.ReplaceAll(companyName, " ", "+")
	pageLink := fmt.Sprintf("%s/search?q=%s", baseURL, fullCompanyName)
	resp, httpErr := http.Get(pageLink) //nolint
	if httpErr != nil {                 //nolint
		err := fmt.Errorf("an error occured trying to scrape google for %s %w", companyName, httpErr)
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err //nolint
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("non-successful response returned with status code: %d", resp.StatusCode)
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

func AboutUs(companyURL string) (string, error) {
	aboutUsURL := ""
	resp, err := http.Get(companyURL)
	if err != nil {
		log.Printf("Error fetching %s: %v\n", companyURL, err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("received non-200 status code: %d", resp.StatusCode)
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return "", err
	}
	re := regexp.MustCompile(`(?i)<a[^>]+href=["']([^"']+)["'][^>]*>(?:\s*about\s*us\s*|about|contact\s*us)\s*</a>`)
	match := re.FindStringSubmatch(string(data))
	if len(match) > 1 {
		aboutUsURL = match[1]
	}
	if aboutUsURL == "" {
		return "", nil
	}
	if !strings.HasPrefix(aboutUsURL, "http") {
		aboutUsURL = companyURL + aboutUsURL
	}
	return aboutUsURL, nil
}

func ExtractEmail(content string) (string, error) {
	resp, err := http.Get(content)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("received non-200 status code: %d", resp.StatusCode)
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	match := re.FindString(string(data))
	return match, nil

}
