# Go functions
A function is a block of statements that can be used repeatedly in a program.

A function will not execute automatically when a page loads.

A function will be executed by a call to the function.

## Create a Function
To create (often referred to as declare) a function, do the following:

Use the func keyword.
Specify a name for the function, followed by parentheses ().
Finally, add code that defines what the function should do, inside curly braces {}.
Syntax
func FunctionName() {
  // code to be executed
}
## Call a Function
Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

In the example below, we create a function named "myMessage()". The opening curly brace ( { ) indicates the beginning of the function code, and the closing curly brace ( } ) indicates the end of the function. The function outputs "I just got executed!". To call the function, just write its name followed by two parentheses ():

Example
```go
package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}
Result:

I just got executed!
A function can be called multiple times.

Example
```
```go
package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage()
  myMessage()
  myMessage()
}
Result:

I just got executed!
I just got executed!
I just got executed!

```
# Naming Rules for Go Functions
A function name must start with a letter
A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
Function names are case-sensitive
A function name cannot contain spaces
If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used
Tip: Give the function a name that reflects what the function does!


# Go Function Parameters and Arguments
Parameters and Arguments
Information can be passed to functions as a parameter. Parameters act as variables inside the function.

Parameters and their types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma:

Syntax
func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
}
Function With Parameter Example
The following example has a function with one parameter (fname) of type string. When the familyName() function is called, we also pass along a name (e.g. Liam), and the name is used inside the function, which outputs several different first names, but an equal last name:

Example
```go
package main
import ("fmt")

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}
Result:

Hello Liam Refsnes
Hello Jenny Refsnes
Hello Anja Refsnes
Note: When a parameter is passed to the function, it is called an argument. So, from the example above: fname is a parameter, while Liam, Jenny and Anja are arguments.
```

# Multiple Parameters
Inside the function, you can add as many parameters as you want:

Example
```go
package main
import ("fmt")

func familyName(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main() {
  familyName("Liam", 3)
  familyName("Jenny", 14)
  familyName("Anja", 30)
}
Result:

Hello 3 year old Liam Refsnes
Hello 14 year old Jenny Refsnes
Hello 30 year old Anja Refsnes
Note: When you are working with multiple parameters, the function call must have the same number of arguments as there are parameters, and the arguments must be passed in the same order.
```

# Go Function Returns
Return Values
If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function:

Syntax
func FunctionName(param1 type, param2 type) type {
  // code to be executed
  return output
}
Function Return Example
Example
```go
Here, myFunction() receives two integers (x and y) and returns their addition (x + y) as integer (int):

package main
import ("fmt")

func myFunction(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(myFunction(1, 2))
}
Result:

3
```
## Named Return Values
In Go, you can name the return values of a function.

Example
```go
Here, we name the return value as result (of type int), and return the value with a naked return (means that we use the return statement without specifying the variable name):

package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return
}

func main() {
  fmt.Println(myFunction(1, 2))
}
Result:

3
The example above can also be written like this. Here, the return statement specifies the variable name:
```
Example
```go
package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return result
}

func main() {
  fmt.Println(myFunction(1, 2))
}
```
## Go Recursion Functions
Recursion Functions
Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.

In the following example, testcount() is a function that calls itself. We use the x variable as the data, which increments with 1 (x + 1) every time we recurse. The recursion ends when the x variable equals to 11 (x == 11). 

Example
```go
package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
}
Result:

1
2
3
4
5
6
7
8
9
10
```
Recursion is a common mathematical and programming concept. This has the benefit of meaning that you can loop through data to reach a result.

The developer should be careful with recursion functions as it can be quite easy to slip into writing a function which never terminates, or one that uses excess amounts of memory or processor power. However, when written correctly recursion can be a very efficient and mathematically-elegant approach to programming.

In the following example, factorial_recursion() is a function that calls itself. We use the x variable as the data, which decrements (-1) every time we recurse. The recursion ends when the condition is not greater than 0 (i.e. when it is 0).

## Example
```go
package main
import ("fmt")

func factorial_recursion(x float64) (y float64) {
  if x > 0 {
     y = x * factorial_recursion(x-1)
  } else {
     y = 1
  }
  return
}

func main() {
  fmt.Println(factorial_recursion(4))
}
Result:

24
```
To a new developer it can take some time to work out how exactly this works, best way to find out is by testing and modifying it.



