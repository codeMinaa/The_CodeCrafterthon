# Go Operators
Operators are used to perform operations on variables and values.

The + operator adds together two values, like in the example below:

* Example:
package main
import ("fmt")

func main() {
  var a = 15 + 25
  fmt.Println(a)
}
Although the + operator is often used to add together two values, it can also be used to add together a variable and a value, or a variable and another variable:

* Example:
package main
import ("fmt")

func main() {
  var (
    sum1 = 100 + 50 // 150 (100 + 50)
    sum2 = sum1 + 250 // 400 (150 + 250)
    sum3 = sum2 + sum2 // 800 (400 + 400)
  )
  fmt.Println(sum3)
}
Go divides the operators into the following groups:

Arithmetic operators
Assignment operators
Comparison operators
Logical operators
Bitwise operators

## Go Arithmetic Operators
Arithmetic Operators
Arithmetic operators are used to perform common mathematical operations.

Operator	Name	Description	Example	Try it
+	Addition	Adds together two values	x + y	
-	Subtraction	Subtracts one value from another	x - y	
*	Multiplication	Multiplies two values	x * y	
/	Division	Divides one value by another	x / y	
%	Modulus	Returns the division remainder	x % y	
++	Increment	Increases the value of a variable by 1	x++	
--	Decrement	Decreases the value of a variable by 1	x--	

