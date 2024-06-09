// You can access a specific slice element by referring to the index number.
package main

import (
	"fmt"
)

func main() {
	myslice := []int{1, 10, 100, 1000}
	fmt.Println(myslice[0])
	fmt.Println(myslice[2])
}
