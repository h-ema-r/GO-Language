// A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
// type Person struct {
// 	name string
// 	age int
// 	job string
// 	salary int
// }

package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	// pers1 specification
	pers1.name = "hege"
	pers1.age = 20
	pers1.job = "teacher"
	pers1.salary = 6000

	// pers1 specification
	pers2.name = "Hema"
	pers2.age = 50
	pers2.job = "Doctor"
	pers2.salary = 100000

	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

}
