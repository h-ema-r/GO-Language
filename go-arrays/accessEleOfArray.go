/* You can access a specific array element by referring to the index number.

In Go, array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.*/

package main

import (
	"fmt"
)

func main() {
	prices := [3]int{1, 2, 3}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
}
