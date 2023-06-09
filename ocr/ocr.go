package ocr

import (
	"strings"
)

var digitTable = map[string]string{
	" _ | ||_|": "0",
	"     |  |": "1",
	" _  _||_ ": "2",
	" _  _| _|": "3",
	"   |_|  |": "4",
	" _ |_  _|": "5",
	" _ |_ |_|": "6",
	" _   |  |": "7",
	" _ |_||_|": "8",
	" _ |_| _|": "9",
}

// Recognize takes a string of digits and returns a slice of strings
func Recognize(s string) []string {

	var out []string

	// Split the string into lines
	lines := strings.Split(s, "\n")

	// Loop through the lines and recognize the digits
	// knowing that each digit is 3 characters wide
	for i := 1; i < len(lines); i = i + 4 {
		out = append(out, recognizeLine(lines[i:i+3]))
	}
	return out
}

func recognizeLine(line []string) string {

	var digitLine string

	// Loop through the line and recognize the digits
	for i := 0; i < len(line[0]); i += 3 {
		digitLine += recognizeDigit(line[0][i:i+3] + line[1][i:i+3] + line[2][i:i+3])
	}
	return digitLine
}

func recognizeDigit(s string) string {
	if digit, ok := digitTable[s]; ok {
		return digit
	}
	return "?"
}
