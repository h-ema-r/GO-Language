/* Append One Slice To Another Slice
To append all the elements of one slice to another slice, use the append()function:

Syntax
slice3 = append(slice1, slice2...)
Note: The '...' after slice2 is necessary when appending the elements of one slice to another.
*/

package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3=%v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n", cap(myslice3))
}
