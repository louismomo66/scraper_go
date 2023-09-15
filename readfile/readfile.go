package readfile

import (
	"bufio"
	"fmt"
	"os"
)

func ReadTxt(path string) ([]string, error) {
	readFile, err := os.Open(path)
	if err != nil {
		err := fmt.Errorf("error occured trying to open file %w", err)
		return nil, err
	}

	defer readFile.Close()

	var fileLines []string
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines, nil
}
