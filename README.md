# A Library for Others

CSV file reader Library. \
Formating based on: 
https://datatracker.ietf.org/doc/html/rfc4180


## Interface 
    type CSVParser interface  {
        ReadLine(r io.Reader) (string, error)
        GetField(n int) (string, error)
        GetNumberOfFields() int
    }

### ReadLine
- Reads one line from open input file. 
- Calling **ReadLine** in a loop allows you to sequentially read each line from the file, continuing until the end of the file is reached.
- Returns line, with terminator removed and and an error if occured.
    - Assumes that input lines are terminated by \r, \n, \r\n, or EOF
- If the line has a missing or extra quote, it should return an empty string and an ErrQuote error.

### GetField
- Returns n-th field from last line read by ReadLine;
- Returns error if n < 0 or beyond last field
- Fields may be surrounded by "..."; such quotes are removed
   
### GetNumberOfFields
- Returns number of fields on last line read by **ReadLine**
- Behavior undefined if called before **ReadLine** is called

## Author
Made on foundation
Daniyar Kabirov ;p