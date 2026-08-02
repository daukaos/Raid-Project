package main

import "fmt"

func printSchedule(rows []string, busy string, important string) {
	if len(rows) == 0 {
		fmt.Println("error: empty schedule")
		return
	}

	expectedLen := len(rows[0])

	for _, row := range rows {
		if len(row) == 0 {
			fmt.Println("error: empty schedule")
			return
		}
		if len(row) != expectedLen {
			fmt.Println("error: rows have different lengths")
			return
		}
		for i := 0; i < len(row); i++ {
			c := row[i]
			if c != '0' && c != '1' && c != '2' {
				fmt.Println("error: invalid symbol (only 0/1/2 allowed)")
				return
			}
		}
	}

}

func main() {
	printSchedule([]string{"0102", "1100", "0010", "2001"}, "###", "!!!")
}
