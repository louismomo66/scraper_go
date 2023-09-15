package main

import (
	"fmt"
	"log"
	"os"
	"scraper/readfile"
	scrape "scraper/scrapers"
)

func main() {
	companyName, _ := readfile.ReadTxt("file.txt")
	file, err := os.Create("results.txt")
	if err != nil { //nolint
		log.Fatal(err)
	}
	defer file.Close()
	for _, name := range companyName { //nolint
		companyUrls, _ := scrape.GetUrls(name)
		aboutUsLink, _ := scrape.AboutUs(companyUrls)
		email, _ := scrape.ExtractEmail(aboutUsLink)
		result := fmt.Sprintf("%s: %s \n", name, email)
		log.Println(result)
		if result != "" { //nolint
			data := []byte(result)
			_, err := file.Write(data)
			if err != nil { //nolint
				log.Printf("Error: %v", err)
			}
		}
	}
}
