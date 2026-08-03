package main

import (
	"fmt"
	"strings"
)

const (
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

// This function is a helper-function
func centerText(s string) string {
	const width = 9
	padding := width - len(s)
	if padding <= 0 {
		return s
	}
	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}

func printSchedule(rows []string, busy string, important string) {

	// This part of the code checks if the given input is valid
	if len(rows) == 0 {
		fmt.Println("error: empty schedule")
		return
	}

	if len(busy) > 9 || len(busy) <= 0 {
		fmt.Println("error: invalid number of characters")
		return
	}
	if len(important) > 9 || len(important) <= 0 {
		fmt.Println("error: invalid number of characters")
		return
	}

	expectedLen := len(rows[0])

	// This part of the code calculates and gives the busiest row
	var maxRowSum, busiestRow int

	for rowIndex, row := range rows {
		if len(row) == 0 {
			fmt.Println("error: empty schedule")
			return
		}
		if len(row) != expectedLen {
			fmt.Println("error: rows have different lengths")
			return
		}

		currentRowSum := 0
		for i := 0; i < len(row); i++ {
			c := row[i]
			if c != '0' && c != '1' && c != '2' {
				fmt.Println("error: invalid symbol (only 0/1/2 allowed)")
				return
			}
			if c != '0' {
				currentRowSum++
			}
		}

		// This part updates the busiest row
		if currentRowSum > maxRowSum {
			maxRowSum = currentRowSum
			busiestRow = rowIndex + 1
		}
	}

	// This part of the code calculates and gives the busiest column
	var maxColSum, busiestCol int

	for col := 0; col < expectedLen; col++ {
		currentColSum := 0
		for row := 0; row < len(rows); row++ {
			if rows[row][col] != '0' {
				currentColSum++
			}
		}

		// This part updates the busiest column
		if currentColSum > maxColSum {
			maxColSum = currentColSum
			busiestCol = col + 1
		}
	}

	// This part prints the table itself
	cellBorder := "+---------"
	border := strings.Repeat(cellBorder, expectedLen) + "+"

	for _, row := range rows {
		fmt.Println(border)
		var sb strings.Builder
		sb.WriteString("|")

		for i := 0; i < len(row); i++ {
			var cellContent string

			switch row[i] {
			case '1':
				cellContent = ColorRed + centerText(busy) + ColorReset
			case '2':
				cellContent = ColorYellow + centerText(important) + ColorReset
			default:
				cellContent = centerText("")
			}

			sb.WriteString(cellContent)
			sb.WriteString("|")
		}
		fmt.Println(sb.String())
	}
	fmt.Println(border)

	// This part prints statistics
	fmt.Println("Busiest column: ", busiestCol)
	fmt.Println("Busiest row: ", busiestRow)
}

func main() {
	printSchedule([]string{"01002", "10111", "00010", "20001"}, "#", "!")
}
