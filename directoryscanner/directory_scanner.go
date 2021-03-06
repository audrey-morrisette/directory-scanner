package directoryscanner

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var compiledRegexes = map[string][]*regexp.Regexp{
	"Credit Card":   {regexp.MustCompile("^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$")},
	"SSN":           {regexp.MustCompile("(^\\d{3}-?\\d{2}-?\\d{4}$|^XXX-XX-XXXX$)")},
	"Word Password": {regexp.MustCompile("password")},
	"Word Username": {regexp.MustCompile("username")},
	"Email":         {regexp.MustCompile("^[\\w\\.=-]+@[\\w\\.-]+\\.[\\w]{2,3}$")},
}

// will hold results from the scan
var results []string

var progress chan int = make(chan int)

func scanFiles(path string, info os.FileInfo, err error) error {

	file, _ := os.Open(path)
	fscanner := bufio.NewScanner(file)
	lineNumber := 1
	// skip the source code
	if file.Name() != "directory_scanner.go" {
		for fscanner.Scan() {
			for key, cr := range compiledRegexes {
				for _, r := range cr {
					if found := r.Find([]byte(fscanner.Text())); found != nil {
						resultsString := key + `,` + string(found) + `,` + file.Name() + `,` + strconv.Itoa(lineNumber)
						results = append(results, resultsString)

					}
				}
			}
			lineNumber++
		}
		progress <- 1
	}

	return nil
}

// Dig ...
func Dig(path string) ([]string, error) {
	err := filepath.Walk(path, scanFiles)
	if err != nil {
		return nil, err
	}
	progress <- 2
	close(progress)

	return results, nil
}
