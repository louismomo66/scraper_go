package main

import (
	"fmt"
	"log"
	"os"
	"scraper/readfile"
	scrape "scraper/scrapers"
)

func main() {
	companyName := readfile.ReadTxt("file.txt")
	file, err := os.Create("results.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, name := range companyName {
		companyUrls := scrape.GetUrls(name)
		aboutUsLink := scrape.AboutUs(companyUrls)
		email := scrape.ExtractEmail(aboutUsLink)
		result := fmt.Sprintf("%s: %s \n", name, email)
		log.Println(result)
		if result != "" {
			data := []byte(result)
			_, err := file.Write(data)
			if err != nil {
				log.Printf("Error: %v", err)
			}
		}
	}
}
