//

package main

import (
	"bufio"
	"fmt"
	"log"
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

func read_txt(path string) []string{
  
  readFile, err := os.Open(path)
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  var fileLines []string
  for fileScanner.Scan() {
    fileLines = append(fileLines,fileScanner.Text())
  }

  readFile.Close()
  return fileLines
}

func getUrls(path2 string) []string {
  names := read_txt(path2)
  var urlResults []string
  for _, companyName := range names {
    doc, err := goquery.NewDocument("http://google.com/search?q="+companyName)
    if err != nil {
      log.Fatal(err)
    }
  
    doc.Find("body a").Each(func(index int, item *goquery.Selection){
  linkTag := item
  link, _:= linkTag.Attr("href")

  if strings.HasPrefix(link,"/url?q=") {
    link = strings.TrimPrefix(link, "/url?q=")
    link = strings.Split(link, "&")[0]
    urlResults = append(urlResults,link)
  }
  
    })
    
}
return urlResults
}


func main() {
links := getUrls("file.txt")

for i,link:= range links {
 fmt.Printf("Links #%d: %s\n", i+1, link)
 }
}
