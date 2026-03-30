# 
Go Constants
If a variable should have a fixed value that cannot be changed, you can use the const keyword.

The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

Syntax
const CONSTNAME type = value
Note: The value of a constant must be assigned when you declare it.

 ## Declaring a Constant
Here is an example of declaring a constant in Go:

*Example
package main
import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
}
* Constant Rules
* Constant names follow the same naming rules as variables
* Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
* Constants can be declared both inside and outside of a function

## Constant Types
There are two types of constants:

Typed constants
Untyped constants
Typed Constants
Typed constants are declared with a defined type:

* Example:
package main
import ("fmt")

const A int = 1

func main() {
  fmt.Println(A)
}

## Untyped Constants
* Untyped constants are declared without a type:

* Example:
package main
import ("fmt")

const A = 1

func main() {
  fmt.Println(A)
}
Note: In this case, the type of the constant is inferred from the value (means the compiler decides the type of the constant, based on the value).

## Constants: Unchangeable and Read-only
When a constant is declared, it is not possible to change the value later:

* Example:
package main
import ("fmt")

func main() {
  const A = 1
  A = 2
  fmt.Println(A)
}
Result:

./prog.go:8:7: cannot assign to A


## Multiple Constants Declaration
Multiple constants can be grouped together into a block for readability:

* Example:
```go
package main
import ("fmt")

const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}
```
