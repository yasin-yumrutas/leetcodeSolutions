package main

import (
	"fmt"
	"strings"
)

// convert converts the string s into a zigzag pattern on numRows and reads line by line.
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]strings.Builder, numRows)

	curRow := 0
	goingDown := false

	for i := 0; i < len(s); i++ {
		rows[curRow].WriteByte(s[i])

		// change direction at the top or bottom row
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	var result strings.Builder
	for i := 0; i < numRows; i++ {
		result.WriteString(rows[i].String())
	}
	return result.String()
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3)) // "PAHNAPLSIIGYIR"
	fmt.Println(convert("PAYPALISHIRING", 4)) // "PINALSIGYAHRPI"
	fmt.Println(convert("A", 1))              // "A"
}
