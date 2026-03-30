# Go Output Functions
Go has three functions to output text:

Print()
Println()
Printf()
The Print() Function
The Print() function prints its arguments with their default format.

* Example:
Print the values of i and j:

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)
}
Result:

*HelloWorld!:

* Example1:
If we want to print the arguments in new lines, we need to use \n.

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, "\n")
  fmt.Print(j, "\n")
}
Result:

Hello
World
Tip: \n creates new lines.

* Example2:
It is also possible to only use one Print() for printing multiple variables.

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, "\n",j)
}
Result:

Hello
World
* Example3:
If we want to add a space between string arguments, we need to use " ":

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, " ", j)
}
Result:

Hello World
* Example4:
Print() inserts a space between the arguments if neither are strings:

package main
import ("fmt")

func main() {
  var i,j = 10,20

  fmt.Print(i,j)
}
Result:

10 20
## The Println() Function
The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:

* Example:
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Println(i,j)
}
Result:

* Hello World!:
## 
The Printf() Function
The Printf() function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

%v is used to print the value of the arguments
%T is used to print the type of the arguments
* Example:
package main
import ("fmt")

func main() {
  var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}
Result:

i has value: Hello and type: string
j has value: 15 and type: int


## Go Formatting Verbs
* Formatting Verbs for Printf()
Go offers several formatting verbs that can be used with the Printf() function.

* General Formatting Verbs
The following verbs can be used with all data types:

Verb	Description
%v	Prints the value in the default format
%#v	Prints the value in Go-syntax format
%T	Prints the type of the value
%%	Prints the % sign
Example
package main
import ("fmt")

func main() {
  var i = 15.5
  var txt = "Hello World!"

  fmt.Printf("%v\n", i)
  fmt.Printf("%#v\n", i)
  fmt.Printf("%v%%\n", i)
  fmt.Printf("%T\n", i)

  fmt.Printf("%v\n", txt)
  fmt.Printf("%#v\n", txt)
  fmt.Printf("%T\n", txt)
}
Result:

15.5
15.5
15.5%
float64
Hello World!
"Hello World!"
string.
Integer Formatting Verbs
The following verbs can be used with the integer data type:

Verb	Description
%b	Base 2
%d	Base 10
%+d	Base 10 and always show sign
%o	Base 8
%O	Base 8, with leading 0o
%x	Base 16, lowercase
%X	Base 16, uppercase
%#x	Base 16, with leading 0x
%4d	Pad with spaces (width 4, right justified)
%-4d	Pad with spaces (width 4, left justified)
%04d	Pad with zeroes (width 4)
* Example:
package main
import ("fmt")

func main() {
  var i = 15
 
  fmt.Printf("%b\n", i)
  fmt.Printf("%d\n", i)
  fmt.Printf("%+d\n", i)
  fmt.Printf("%o\n", i)
  fmt.Printf("%O\n", i)
  fmt.Printf("%x\n", i)
  fmt.Printf("%X\n", i)
  fmt.Printf("%#x\n", i)
  fmt.Printf("%4d\n", i)
  fmt.Printf("%-4d\n", i)
  fmt.Printf("%04d\n", i)
}
Result:

1111
15
+15
17
0o17
f
F
0xf
  15
15
0015.

## String Formatting Verbs
The following verbs can be used with the string data type:

Verb	Description
%s	Prints the value as plain string
%q	Prints the value as a double-quoted string
%8s	Prints the value as plain string (width 8, right justified)
%-8s	Prints the value as plain string (width 8, left justified)
%x	Prints the value as hex dump of byte values
% x	Prints the value as hex dump with spaces
* Example:
package main
import ("fmt")

func main() {
  var txt = "Hello"
 
  fmt.Printf("%s\n", txt)
  fmt.Printf("%q\n", txt)
  fmt.Printf("%8s\n", txt)
  fmt.Printf("%-8s\n", txt)
  fmt.Printf("%x\n", txt)
  fmt.Printf("% x\n", txt)
}
Result:

Hello
"Hello"
   Hello
Hello
48656c6c6f
48 65 6c 6c 6f

## Boolean Formatting Verbs
The following verb can be used with the boolean data type:

Verb	Description
%t	Value of the boolean operator in true or false format (same as using %v)
* Example:

package main
import ("fmt")

func main() {
  var i = true
  var j = false

  fmt.Printf("%t\n", i)
  fmt.Printf("%t\n", j)
}
Result:

true
false

## Float Formatting Verbs
The following verbs can be used with the float data type:

Verb	Description
%e	Scientific notation with 'e' as exponent
%f	Decimal point, no exponent
%.2f	Default width, precision 2
%6.2f	Width 6, precision 2
%g	Exponent as needed, only necessary digits
* Example:
```go
package main
import ("fmt")

func main() {
  var i = 3.141

  fmt.Printf("%e\n", i)
  fmt.Printf("%f\n", i)
  fmt.Printf("%.2f\n", i)
  fmt.Printf("%6.2f\n", i)
  fmt.Printf("%g\n", i)
}
Result:

3.141000e+00
3.141000
3.14
  3.14
3.141
```



