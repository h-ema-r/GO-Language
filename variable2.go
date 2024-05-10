//In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type

package main

import (
	"fmt"
)

func main() {
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
