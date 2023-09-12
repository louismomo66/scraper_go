package readfile_test

import (
	"os"
	"scraper/readfile"
	"testing"
)

func TestReadTxt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			"multiplelines",
			"codebits\ncodeclinic\njumia\n",
			[]string{"codebits", "codeclinic", "jumia"},
		},
		{
			"singleline",
			"foodhub",
			[]string{"foodhub"},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tempFile, err := os.CreateTemp("", "testfile.txt")

			if err != nil {
				t.Fatalf("Failed to create temporary file %v", err)
			}

			defer os.Remove(tempFile.Name())

			defer tempFile.Close()

			_, err = tempFile.WriteString(tc.input)
			if err != nil {
				t.Fatalf("Failed to write test data to temporary test file: %v", err)
			}

			result := readfile.ReadTxt(tempFile.Name())
			if len(result) != len(tc.want) {
				t.Errorf("Expected %d lines,but got %d lines", len(tc.want), len(result))
			}

			for i := range tc.want {
				if result[i] != tc.want[i] {
					t.Errorf("%s: Expected line %d to be \"%s\", but got \"%s\"", tc.name, i+1, tc.want[i], result[i])
				}
			}
		})
	}
}
