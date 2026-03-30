# Go Comments
A comment is a text that is ignored upon execution.

Comments can be used to explain the code, and to make it more readable.

Comments can also be used to prevent code execution when testing an alternative code.

Go supports single-line or multi-line comments.
## Go Single-line Comments
Single-line comments start with two forward slashes (//).

Any text between // and the end of the line is ignored by the compiler (will not be executed).

* Example1:
// This is a comment
package main
import ("fmt")

func main() {
  // This is a comment
  fmt.Println("Hello World!")
}
* Note: The following example uses a single-line comment at the end of a code line:

* Example2: 
Example
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!") // This is a comment
}

## Go Multi-line Comments
Multi-line comments start with /* and ends with */.

Any text between /* and */ will be ignored by the compiler:

* Example: 

Example
package main
import ("fmt")

func main() {
  /* The code below will print Hello World
  to the screen, and it is amazing */
  fmt.Println("Hello World!")
}
## Comment to Prevent Code Execution
You can also use comments to prevent the code from being executed.

The commented code can be saved for later reference and troubleshooting.

* Example: 
```go
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
  // fmt.Println("This line does not execute")
}```go