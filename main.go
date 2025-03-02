package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

var (
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

type YourCSVParser struct {
	fields []string
}

// func Split(text string) ([]string, error) {

// }

func (parser *YourCSVParser) ReadLine(r io.Reader) (string, error) {
	char := make([]byte, 1)
	text := ""
	var err error
	array := make([]string, 0)
	quotesFlag := false

	element := ""

	for {
		_, err = r.Read(char)
		if char[0] == '\n' || char[0] == '\r' || err != nil {
			break
		}
		text += string(char[0])
	}
	// len 10
	// 0 1 2 3 4 5 6C 7R 8L 9F // i = 6
	// +3 = 9
	for i, a := range text {
		if len(element) == 0 && a == '"' { // quoteflag start
			quotesFlag = true
			continue
		}

		if !quotesFlag && a == 'C' { // CRLF check
			if i+3 < len(text) {
				if text[i+1] == 'R' && text[i+2] == 'L' && text[i+3] == 'F' {
					break
				}
			}
		}

		if quotesFlag && a == '"' { // quote and quoteflag end
			if i+1 < len(text) {
				if text[i+1] == '"' {
					i++
					element += string('"')
				} else {
					quotesFlag = false
					continue
				}
			} else {
				quotesFlag = false
				continue
			}
		}

		if !quotesFlag && a == ',' {
			array = append(array, element)
			element = ""
			continue
		}

		element += string(a)

	}
	if len(strings.TrimSpace(element)) != 0 {
		array = append(array, element)
	}
	if quotesFlag == true {
		err = ErrQuote
	}
	parser.fields = array
	return text, err
}

func (parser *YourCSVParser) GetField(n int) (string, error) {
	// fmt.Println(parser.GetNumberOfFields())
	if n < 0 || n >= parser.GetNumberOfFields() {
		return "", ErrFieldCount
	}
	return parser.fields[n], nil
}

func (parser *YourCSVParser) GetNumberOfFields() int {
	return len(parser.fields)
}

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var csvparser CSVParser = &YourCSVParser{}

	for {
		line, err := csvparser.ReadLine(file)
		fmt.Println(line)
		fmt.Println(csvparser.GetNumberOfFields())
		field, _ := csvparser.GetField(2)
		fmt.Println(field)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}
	}
}
