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

func GetUrls(companyName string) string {
	resp, err := http.Get("http://google.com/search?q=" + companyName)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
	}
	foundURL := ""
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
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
	
		urlResults := foundURL

return urlResults
}

func ExtractEmail(content string) string { 

	resp,err := http.Get(content)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	data,err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	
	re := regexp.MustCompile(`[\w\.-]+@[\w\.-]+`)
	match := re.FindString(string(data))
	return match
}

func AboutUs(companyURL string) string {
	aboutUsURL := ""
	resp, err := http.Get(companyURL)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", companyURL, err)

	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading respinse body: %v\n", err)

	}

	re := regexp.MustCompile(`(?i)<a[^>]+href=["']?([^"']+)["']?[^>]*>(?:\s*about\s*us\s*|about|contact\s*us)\s*</a>`)
	match := re.FindStringSubmatch(string(data))
	if len(match) > 1 {
		aboutUsURL = match[1]
	}
	if !strings.HasPrefix(aboutUsURL, "http") {
		aboutUsURL = companyURL + aboutUsURL
	}
	return aboutUsURL
}
