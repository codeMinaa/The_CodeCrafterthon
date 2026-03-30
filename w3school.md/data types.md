# Go Data Types
Data type is an important concept in programming. Data type specifies the size and type of variable values.

Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.

Go has three basic data types:

bool: represents a boolean value and is either true or false
Numeric: represents integer types, floating point values, and complex types
string: represents a string value
* Example:
This example shows some of the different data types in Go:

package main
import ("fmt")

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)
}


## Go Boolean Data Type
* Boolean Data Type:
A boolean data type is declared with the bool keyword and can only take the values true or false.

The default value of a boolean data type is false.

* Example:
This example shows some different ways to declare Boolean variables:

package main
import ("fmt")

func main() {
  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value
  b4 := true // untyped declaration with initial value

  fmt.Println(b1) // Returns true
  fmt.Println(b2) // Returns true
  fmt.Println(b3) // Returns false
  fmt.Println(b4) // Returns true
}


## Go Integer Data Types
Integer data types are used to store a whole number without decimals, like 35, -50, or 1345000.

The integer data type has two categories:

Signed integers - can store both positive and negative values
Unsigned integers - can only store non-negative values
Tip: The default type for integer is int. If you do not specify a type, the type will be int.


* Signed Integers:
Signed integers, declared with one of the int keywords, can store both positive and negative values:

Example
package main
import ("fmt")

func main() {
  var x int = 500
  var y int = -4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}
* package main
import ("fmt")

func main() {
  var x uint = 500
  var y uint = 4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}

## Go Float Data Types
The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.
* Tip: The default type for float is float64. If you do not specify a type, the type will be float64.

## The float32 Keyword
Example
This example shows how to declare some variables of type float32:

package main
import ("fmt")

func main() {
  var x float32 = 123.78
  var y float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}

## The float64 Keyword
The float64 data type can store a larger set of numbers than float32.

Example
This example shows how to declare a variable of type float64:

package main
import ("fmt")

func main() {
  var x float64 = 1.7e+308
  fmt.Printf("Type: %T, value: %v", x, x)
}

## Which Float Type to Use?
The type of float to choose, depends on the value the variable has to store.

Example
This example will result in an error because 3.4e+39 is out of range for float32:

package main
import ("fmt")

func main() {
  var x float32= 3.4e+39
  fmt.Println(x)
}
Result:

./prog.go:5:7: constant 3.4e+39 overflows float32

## Go String Data Type
String Data Type
The string data type is used to store a sequence of characters (text). String values must be surrounded by double quotes:

Example
package main
```go
import ("fmt")

func main() {
  var txt1 string = "Hello!"
  var txt2 string
  txt3 := "World 1"

  fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
  fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
  fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}
Result:

Type: string, value: Hello!
Type: string, value:
Type: string, value: World 1
```