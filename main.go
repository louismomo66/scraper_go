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

func extractEmail(content string) string { 

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

func aboutUs(companyName, companyURL string) string {
	aboutUsURL := ""
	// value := ""
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

func main() {
	companyName := readTxt("file.txt")
	companyUrls := getUrls("file.txt")
// emailC := extractEmail("https://sunbird.ai/about/")
// fmt.Println(emailC)
	file, err := os.Create("results.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for i, comp := range companyUrls {

		if i < len(companyName) {
			returned := aboutUs(companyName[i], comp)
			data1 := []byte(returned + "\n")
			// temp := filepath.Join(os.TempDir(),"results.txt")
			_, err := file.Write(data1)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(returned)
		}
	}
}
