# directory-scanner
This is a simple directory scanner written in the Go programming language. It identifies sensitive information in files located in the specified directory, and returns the file name that contains the sensitive data, the kind of sensitive data it discovered, and the location of the sensitive data in the file

### To Use:
`go run directory-scanner.go [directory] [save location]`

### Currently Detects:
- Credit Cards
- Social Security Numbers
- ~~AWS Access Keys~~
- Strings "username" and "password"

## TODO: 
- [x] Create output report
- [x] Add ability to choose location to save report
- [ ] Add improved console output
- [ ] Add output options (.csv, .tsv, etc.)
- [ ] Enable to check zipped files/more file types
- [ ] Add more items to detect