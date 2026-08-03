Schedule Table

A small Go program that prints a weekly/daily schedule as a colored ASCII table.

To run it write in terminal "go run main.go"
The rules:
Each character in a row represents one time slot:

Character	Meaning	Color
0	Free	none
1	Busy	red 
2	Important	yellow 

Example:
func main() {
	printSchedule([]string{"01002", "10100", "00010", "20001"}, "#", "!")
}
When this code runs it gives this:
+---------+---------+---------+---------+---------+
|         |    #    |         |         |    !    |
+---------+---------+---------+---------+---------+
|    #    |         |    #    |         |         |
+---------+---------+---------+---------+---------+
|         |         |         |    #    |         |
+---------+---------+---------+---------+---------+
|    !    |         |         |         |    #    |
+---------+---------+---------+---------+---------+
(# = red color, ! = yellow color)

Team:
daukaos(validation + color)
kukumberversion2(table + bonus)
yerali67(table + readme.md + bonus)