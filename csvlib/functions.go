package csvlib

import (
	"io"
	"log"
)

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

	for i := 0; i < len(text); i++ {
		if len(element) == 0 && text[i] == '"' { // quoteflag start
			quotesFlag = true
			continue
		}

		if !quotesFlag && text[i] == '"' {
			err = ErrQuote
			return "", err
		}

		if !quotesFlag && text[i] == 'C' { // CRLF check
			if i+3 < len(text) {
				if text[i+1] == 'R' && text[i+2] == 'L' && text[i+3] == 'F' {
					break
				}
			}
		}

		if quotesFlag && text[i] == '"' { // quote and quoteflag end
			if i+1 < len(text) {
				if text[i+1] == '"' {
					i++
				} else {
					quotesFlag = false
					continue
				}
			} else {
				quotesFlag = false
				continue
			}
		}

		if !quotesFlag && text[i] == ',' {
			array = append(array, element)
			element = ""
			continue
		}

		element += string(text[i])

	}
	if len(element) != 0 { // strings.TrimSpace(element)
		array = append(array, element)
	}
	if parser.len == 0 {
		parser.len = len(array)
	}
	if len(array) != parser.len {
		err = ErrFieldCount
	}
	if quotesFlag == true {
		err = ErrQuote
	}
	parser.fields = array
	parser.text = text
	return parser.text, err
}

func (parser *YourCSVParser) GetField(n int) (string, error) {
	// fmt.Println(parser.GetNumberOfFields())
	if n < 0 || n >= parser.GetNumberOfFields() {
		return "", ErrFieldCount
	}
	return parser.fields[n], nil
}

func (parser *YourCSVParser) GetNumberOfFields() int {
	if len(parser.text) == 0 {
		log.Fatal("Undefined behavior: called before ReadLine is called")
	}
	return len(parser.fields)
}
