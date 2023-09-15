package readfile_test

import (
	"errors"
	"fmt"
	"scraper/readfile"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadTxt(t *testing.T) {
	t.Parallel()
	type args struct {
		fileName string
	}
	tests := []struct {
		name      string
		input     args
		want      []string
		wantedErr error
	}{
		{
			"nonexistent",
			// "codebits\ncodeclinic\njumia\n",
			args{
				fileName: "/home/louis/Desktop/scrape/file1.txt",
			},
			[]string{},
			errors.New("error occured trying to open file open /home/louis/Desktop/scrape/file1.txt: no such file or directory"),
		},
		{
			"existentfile",
			args{
				fileName: "../file.txt",
			},
			[]string{"innovex",
				"codebits",
				"codeclinic",
				"jumia",
				"netlabs",
				"sunbirdAi"},
			nil,
		},
	}
	for i := range tests {
		i := i
		t.Run(tests[i].name, func(t *testing.T) {
			t.Parallel()

			result, err := readfile.ReadTxt(tests[i].input.fileName)
			if err != nil && tests[i].wantedErr == nil {
				assert.Fail(t, fmt.Sprintf("Error not expected but got one:\n"+"error: %q", err))
				return
			}
			if tests[i].wantedErr != nil {
				assert.EqualError(t, err, tests[i].wantedErr.Error())
				return
			}
			assert.Equal(t, tests[i].want, result)
		})
	}
}
