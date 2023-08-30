package readfile

import (
	"bufio"
	"log"
	"os"
)
func ReadTxt(path string) []string {

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