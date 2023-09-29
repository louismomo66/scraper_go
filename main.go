package main

import (
	"fmt"
	"log"
	"os"
	"scraper/readfile"
	scrape "scraper/scrapers"
)

func main() {
	const baseURL = "http://google.com/search?q="
	companyName, _ := readfile.ReadTxt("file.txt")
	file, err := os.Create("results.txt")
	if err != nil { //nolint
		log.Fatal(err)
	}
	defer file.Close()

	for _, name := range companyName { //nolint
		companyUrls, err := scrape.GetUrls(baseURL, name)
		aboutUsLink, err := scrape.AboutUs(companyUrls)
		if err != nil {
			log.Printf("Error fetching about us link for %s: %v", companyUrls, err)
			continue
		}
		email, err := scrape.ExtractEmail(aboutUsLink)
		if err != nil {
			log.Printf("Error extracring email from %s: %v", aboutUsLink, err)
			continue
		}
		result := fmt.Sprintf("%s: %s \n", name, email)
		log.Println(result)
		if result != "" { //nolint
			data := []byte(result)
			_, err := file.Write(data)
			if err != nil { //nolint
				log.Printf("Error writing to file: %v", err)
			}
		}
	}
}
