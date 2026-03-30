# Go Variables
Variables are containers for storing data values.

## Go Variable Types
In Go, there are different types of variables, for example:

* int- stores integers (whole numbers), such as 123 or -123
* float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
* string - stores text, such as "Hello World". String values are surrounded by double quotes
* bool- stores values with two states: true or false

## Declaring (Creating) Variables
In Go, there are two ways to declare a variable:
* With the var keyword:
Use the var keyword, followed by variable name and type:

* Syntax
var variablename type = value
Note: You always have to specify either type or value (or both).

* With the := sign:
Use the := sign, followed by the variable value:

* Syntax
variablename := value
Note: In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

Note: It is not possible to declare a variable using :=, without assigning a value to it.

* Variable Declaration With Initial Value
If the value of a variable is known from the start, you can declare the variable and assign a value to it on one line:

* Example::
package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}
Note: The variable types of student2 and x is inferred from their values.

## Variable Declaration Without Initial Value
In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:

* Example:
package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
* Example explained
In this example there are 3 variables:

a
b
c
These variables are declared but they have not been assigned initial values.

By running the code, we can see that they already have the default values of their respective types:

a is ""
b is 0
c is false.

## Value Assignment After Declaration
It is possible to assign a value to a variable after it is declared. This is helpful for cases the value is not initially known.

* Example:
package main
import ("fmt")

func main() {
  var student1 string
  student1 = "John"
  fmt.Println(student1)
}
Note: It is not possible to declare a variable using ":=" without assigning a value to it.

## Difference Between var and :=
There are some small differences between the var var :=:

var	:=
Can be used inside and outside of functions	Can only be used inside functions
Variable declaration and value assignment can be done separately	Variable declaration and value assignment cannot be done separately (must be done in the same line)
* Example:
This example shows declaring variables outside of a function, with the var keyword:
```go
package main
import ("fmt")

var a int
var b int = 2
var c = 3

func main() {
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
* Example:
Since := is used outside of a function, running the program results in an error.

package main
import ("fmt")

a := 1

func main() {
  fmt.Println(a)
}
Result:

./prog.go:5:1: syntax error: non-declaration statement outside function body
 ```