package readfile_test

import (
	"os"
	"scraper/readfile"
	"testing"
)

func TestReadTxt(t *testing.T) {
	tempFile,err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file %v",err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	testData := "good\n bad\n ugly"
	_,err = tempFile.WriteString(testData)
	if err != nil {
		t.Fatalf("Failed to write test data to temporary test file: %v", err)
	}

	result := readfile.ReadTxt(tempFile.Name())
	expected := []string{"good \nbad\nugly"}
	if len(result) != len(expected) {
		t.Errorf("Expected %d lines,but got %d lines", len(expected),len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected line %d to be \"%s\", but got \"%s\"",i+1,expected[i],result[i])
		}
	} 
}