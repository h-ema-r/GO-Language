//The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
