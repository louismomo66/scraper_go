//

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// func readData(filepath string) *os.File {
//   readFile, err := os.Open(filepath)
//   if err != nil {
//     fmt.Println(err)
//   }
//   return readFile}

func readTxt(path string) []string {

	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
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

func getUrls(companyName string) string {
	// names := readTxt(companyName)
	// var urlResults []string
	// for _, companyName := range names {
	url := "http://google.com/search?q=" + companyName
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// doc, err := goquery.NewDocument("http://google.com/search?q=" + companyName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	firstLink := ""
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")

		if strings.HasPrefix(link, "/url?q=") && firstLink == "" {
			link = strings.TrimPrefix(link, "/url?q=")
			link = strings.Split(link, "&")[0]
			firstLink = link
		}

	})

	// }
	return firstLink
}

func main() {
	companyNames := readTxt("file.txt")
	var companyUrls []string

	for _, companyName := range companyNames {
		firstLink := getUrls(companyName)
		companyUrls = append(companyUrls, firstLink)
	}

	for i, link := range companyUrls {
		fmt.Printf("Links #%d: %s\n", i+1, link)
	}
}
