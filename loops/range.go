//The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.

package main

import (
	"fmt"
)

func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Println(idx, val)
	}
}
