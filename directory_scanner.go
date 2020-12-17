package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

var compiledRegexes = map[string][]*regexp.Regexp{
	"Credit Card":   {regexp.MustCompile("^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$")},
	"SSN":           {regexp.MustCompile("(^\\d{3}-?\\d{2}-?\\d{4}$|^XXX-XX-XXXX$)")},
	"Word Password": {regexp.MustCompile("password")},
	"Word Username": {regexp.MustCompile("username")},
	//"AWS Secret Key": {regexp.MustCompile("(?<![A-Z0-9])[A-Z0-9]{20}(?![A-Z0-9])"), regexp.MustCompile("(?<![A-Za-z0-9/+=])[A-Za-z0-9/+=]{40}(?![A-Za-z0-9/+=])")},
}

// will hold results from the scan
var results []string

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
						fmt.Println(resultsString)
						results = append(results, resultsString)
					}
				}
			}
			lineNumber++
		}
	}

	return nil
}

// Dig ...
func Dig(path string) {
	err := filepath.Walk(path, scanFiles)
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		//panic("No directory provided")
		fmt.Println("Program Use: go directory_scanner.go [directory] [output/directory]")
		return
	}
	if os.Args[1] == "help" {
		fmt.Println("-- Program Use: go directory_scanner.go [directory] [output/directory]")
		fmt.Println("-- Use '.' as the argument for directory to scan from the current folder")
		fmt.Println("-- If no location is provided to save the results, the results will automatically be stored in the current directory")
		return
	}

	Dig(os.Args[1])

	var saveDirectory string
	if len(os.Args) == 3 {
		saveDirectory = os.Args[2] + "\\"
	}
	fmt.Println(saveDirectory + "results-" + time.Now().Format("01-02-2006") + ".csv")
	resultsFile, err := os.Create(saveDirectory + "results-" + time.Now().Format("01-02-2006") + ".csv")
	defer resultsFile.Close()
	if err != nil {
		panic(err)
	}
	// writing column headers
	_, err = resultsFile.WriteString("Type,Value,File Name,Location\n")
	if err != nil {
		panic(err)
	}
	for _, result := range results {
		_, err := resultsFile.WriteString(result + "\n")
		if err != nil {
			panic(err)
		}
	}

}
