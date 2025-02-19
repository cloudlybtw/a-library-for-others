package main

import (
	"errors"
	"fmt"
	"io"
	"os"
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

func (parser YourCSVParser) ReadLine(r io.Reader) (string, error) {
	char := make([]byte, 1)
	text := ""
	var err error
	for {
		_, err = r.Read(char)
		if char[0] == '\n' || err != nil {
			break
		}
		text += string(char[0])
	}
	return text, err
}

func (parser YourCSVParser) GetField(n int) (string, error) {
	return "hello", nil
}

func (parser YourCSVParser) GetNumberOfFields() int {
	return 0
}

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var csvparser CSVParser = YourCSVParser{}

	for {
		line, err := csvparser.ReadLine(file)
		fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}
	}
}
