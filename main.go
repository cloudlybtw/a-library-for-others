package main

import (
	"fmt"
	"io"
	"os"

	"a-library-for-others/csvlib"
)

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var csvparser csvlib.CSVParser = &csvlib.YourCSVParser{}

	for {
		line, err := csvparser.ReadLine(file)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}
		fmt.Println(line)
		fmt.Println(csvparser.GetNumberOfFields())
		field, _ := csvparser.GetField(2)
		fmt.Println(field)
	}
}
