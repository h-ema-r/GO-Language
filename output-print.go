// If we want to print the arguments in new lines, we need to use \n.
package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World"
	var x, y = 10, 20

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
	fmt.Print(i, " ", j, "\n")

	// Print() inserts a space between the arguments if neither are strings:
	fmt.Print(x, y)
}
