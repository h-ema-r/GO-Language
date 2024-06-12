// It is possible to have multiple values for each case in the switch statement:

package main

import (
	"fmt"
)

func main() {
	day := 5
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd Weekday")

	case 2, 4:
		fmt.Println("Even Weekday")

	case 6, 7:
		fmt.Println("Weekend")

	default:
		fmt.Println("Invalid day of day number")
	}
}
