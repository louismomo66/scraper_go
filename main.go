package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func readTxt(path string) []string {

	readFile, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}

	defer readFile.Close()

	var fileLines []string
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}

func getUrls(companyName string) []string {
	names := readTxt(companyName)
	var urlResults []string
	for _, companyName := range names {
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
		if foundURL != "" {
			urlResults = append(urlResults, foundURL)
		}
	}
	return urlResults
}


func aboutUs(companyURL string) {
	aboutUsURL := ""
	resp, err := http.Get(companyURL)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", companyURL, err)
		// return aboutUsURL
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading respinse body: %v\n", err)
		// return aboutUsURL
	}

	re := regexp.MustCompile(`(?i)<a[^>]+href=["']?([^"']+)["']?[^>]*>(?:\s*about\s*us\s*|about|contact\s*us)\s*</a>`)
	match := re.FindStringSubmatch(string(data))
	if len(match) > 1 {
		aboutUsURL = match[1]
	}
	if !strings.HasPrefix(aboutUsURL, "http") {
		aboutUsURL = companyURL + aboutUsURL
	}

	fmt.Printf("About us:%s \n", aboutUsURL)
}

func main() {

	companyUrls := getUrls("file.txt")
	for _, comp := range companyUrls {
		// fmt.Println(comp)
		// homeUrl(comp)
		aboutUs(comp)
	}

}
