package main

import (
	"fmt"
)

func main() {
	myslice := make([]int, 5, 10)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)
}
