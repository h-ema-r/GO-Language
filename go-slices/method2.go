// Create a Slice From an Array

package main

import (
	"fmt"
)

func main() {
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)
}
