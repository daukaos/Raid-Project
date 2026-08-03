package main

import (
	"fmt"
)

func printSchedule(rows []string, busy string, important string) {
	if len(rows) == 0 {
		fmt.Println("error: empty schedule")
		return // this part checks if there are no rows to make the schedule,
		//if there are no rows then returns "error:empty schedule"
	}

	if len(busy) > 9 || len(busy) <= 0 {
		fmt.Println("error: invalid number of characters")
		return
	}
	if len(important) > 9 || len(important) <= 0 {
		fmt.Println("error: invalid number of characters")
		return
	}

	expectedLen := len(rows[0]) // it is the variable,it keeps the value of the first row,
	// with this we can compare the lengths of the strings

	for _, row := range rows { // this is the loop, it checks the every given row
		if len(row) == 0 {
			fmt.Println("error: empty schedule")
			return // this part checks every row, if there is one with length of 0 then returns error
		}
		if len(row) != expectedLen {
			fmt.Println("error: rows have different lengths")
			return // this checks the first row to other rows, if the length are different, returns error
		}
		for i := 0; i < len(row); i++ {
			c := row[i]
			if c != '0' && c != '1' && c != '2' {
				fmt.Println("error: invalid symbol (only 0/1/2 allowed)")
				return // this nested loop checks if the characters in every row are the ones that are allowed,
				// if there are characters that are not allowed, returns error
			}
		}
	}

}

func main() {
	printSchedule([]string{"0102", "1100", "0010", "2001"}, "#", "!")
}
