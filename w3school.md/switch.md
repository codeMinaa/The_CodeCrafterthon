# The switch Statement
Use the switch statement to select one of many code blocks to be executed.

The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a break statement.

Single-Case switch Syntax
Syntax
switch expression {
case x:
   // code block
case y:
   // code block
case z:
...
default:
   // code block
}
This is how it works:

The expression is evaluated once
The value of the switch expression is compared with the values of each case
If there is a match, the associated block of code is executed
The default keyword is optional. It specifies some code to run if there is no case match
## Single-Case switch Example
The example below uses a weekday number to calculate the weekday name:

Example
```go
package main
import ("fmt")

func main() {
  day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  }
}
Result:

Thursday
```
