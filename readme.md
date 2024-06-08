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
'''
var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred

'''

- - with := sign language
'''
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred
'''
**Note:** The length specifies the number of elements to store in the array. In Go, arrays have a fixed length. The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values).