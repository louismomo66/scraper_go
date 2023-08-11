//

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func readTxt(path string) []string {

	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
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
			fmt.Println(err)
		}

		doc.Find("body a").Each(func(index int, item *goquery.Selection) {
			linkTag := item
			link, _ := linkTag.Attr("href")

			if strings.HasPrefix(link, "/url?q=") {
				link = strings.TrimPrefix(link, "/url?q=")
				link = strings.Split(link, "&")[0]
				urlResults = append(urlResults, link)
			}

		})

	}
	return urlResults
}

func facebookURL(urls []string) []string {
	var facebookUrls []string
	re := regexp.MustCompile(`^(https?://)?(www\.)?facebook\.com/[^/]+/?$`)

	for _, url := range urls {
		if re.MatchString(url) {
			facebookUrls = append(facebookUrls, url)
		}
	}
	return facebookUrls
}

func main() {

	companyUrls := getUrls("file.txt")
	facebookUrls := facebookURL(companyUrls)
	for i, link := range facebookUrls {
		fmt.Printf("Links #%d: %s\n", i+1, link)
	}
}
