# Go Language

Go is a statically compiled high-level programming language. It is commonly used to create computer programs.

## Introduction

Go, also known as Golang, is a powerful programming language developed by Google. It is designed for simplicity, efficiency, and ease of use. With its built-in concurrency support and robust standard library, Go is well-suited for developing a wide range of applications, from web servers to system utilities.

## Command to run go program
```
go run main.py

```

## Go Comments
Go supports single-line or multi-line comments.These will be ignored by compiler.
- Single-line comments start with two forward slashes (//).
- Multi-line comments start with /* and ends with */.

## Go Variables Types

In Go, variables can hold different types of data. Some of the common variable types include:

- `int`: Represents integer numbers.
- `float32`: Represents floating-point numbers.
- `string`: Represents a sequence of characters.
- `bool`: Represents boolean values (true or false).

### Two ways to declare a variable
- with 'var' keyword:
```
var variablename type = value

```
- with := sign:
```
variablename := value

```

### Go has three functions to output text:
- Print()
- Println()
- Printf()

#### The Printf() Function
The Printf() function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

- %v is used to print the value of the arguments
- %T is used to print the type of the arguments
- %%	Prints the % sign


## Go has three basic data types:

- bool: represents a boolean value and is either true or false
- Numeric: represents integer types, floating point values, and complex types
- string: represents a string value


## Go Arrays
- Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

- Declare an Array
- - with the **var** keyword

```
var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred

```

- - with := sign language
```
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred

```

**Note:** The length specifies the number of elements to store in the array. In Go, arrays have a fixed length. The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values).

## Go Slices:
In Go, there are several ways to create a slice:

- Using the []datatype{values} format
```
slice_name := []datatype{values}
e.g 
myslice := []int{}

```

- Create a slice from an array
```
var myarray = [length]datatype{values} // An array
myslice := myarray[start:end] // A slice made from the array

```

- Using the make() function
```
slice_name := make([]type, length, capacity)
// Note: If the capacity parameter is not defined, it will be equal to length.

```


**In Go, there are two functions that can be used to return the length and capacity of a slice:**

**- len() function** - returns the length of the slice (the number of elements in the slice)<br>
**- cap() function** - returns the capacity of the slice (the number of elements the slice can grow or shrink to)


## Operators

### 1. Airthmetic Operators
 '+',  '-',  '*'  ,'/'  ,'%',  '++'  ,'--'

 ### 2. Assignment Operators
 '=', '+=', '-=', '*=', '/=', '%=', '&=', '|=', '^=', '>>=', '<<='

 ### 3. Comparison Operator
  '==', '!=', '>', '<', '>=', '<='

  ### 4. Logical Operator
  '&&', '||', '!'

  ### 5. Go-Bitwise Operators
  '&', '|', '^', '<<', '>>'

  ## Go Conditions
  - if 
  - else
  - else if 
  - switch
  <br>
 **NOTE:** 
 - Use the switch statement to select one of many code blocks to be executed.

- The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a break statement.


## Go Struct
- A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

- While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

- A struct can be useful for grouping data together to create records.

- Declare a struct 
 To declare a structure in Go, use the type and struct keywords:


 ```
type Person struct {
  name string
  age int
  job string
  salary int
}
```

## Go Maps
- Maps are used to store data values in key:value pairs.
- Each element in a map is a key:value pair.
- A map is an unordered and changeable collection that does not allow duplicates.
- The length of a map is the number of its elements. You can find it using the len() function.
- Syntax
- 1. Create Maps Using var and :=
 ``` 
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}
```
- 2. Create Maps Using make()Function:
 
 ```
var a = make(map[KeyType]ValueType)
b := make(map[KeyType]ValueType)
```
 
 
