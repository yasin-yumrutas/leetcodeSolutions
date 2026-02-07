package main

import "fmt"

// generate returns the first numRows of Pascal's Triangle.
func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		// row length is i+1
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1

		// fill middle elements
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}

		result[i] = row
	}

	return result
}

func main() {
	numRows := 5
	triangle := generate(numRows)

	fmt.Printf("Pascal's Triangle (%d rows):\n", numRows)
	for _, row := range triangle {
		fmt.Println(row)
	}
}
