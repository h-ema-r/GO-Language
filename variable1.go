//Variable Declaration With Initial Value

package main

import (
	"fmt"
)

func main() {
	var student1 string = "Siya" // type is string
	var student2 = "Giya"        // type is inferred
	x := 2                       // type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}
