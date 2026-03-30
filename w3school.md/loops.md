Go For Loops
The for loop loops through a block of code a specified number of times.

The for loop is the only loop available in Go.

Go for Loop
Loops are handy if you want to run the same code over and over again, each time with a different value.

Each execution of a loop is called an iteration.

The for loop can take up to three statements:

Syntax
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}
statement1 Initializes the loop counter value.

statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.

statement3 Increases the loop counter value.

Note: These statements don't need to be present as loops arguments. However, they need to be present in the code in some form.

for Loop Examples
Example 1
```go
This example will print the numbers from 0 to 4:  

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}
Result:

0
1
2
3
4
```
Example 1 explained
i:=0; - Initialize the loop counter (i), and set the start value to 0
i < 5; - Continue the loop as long as i is less than 5
i++ - Increase the loop counter value by 1 for each iteration
## Example 2
This example counts to 100 by tens: 
```go
package main
import ("fmt")

func main() {
  for i:=0; i <= 100; i+=10 {
    fmt.Println(i)
  }
}
Result:

0
10
20
30
40
50
60
70
80
90
100
```
Example 2 explained
i:=0; - Initialize the loop counter (i), and set the start value to 0
i <= 100; - Continue the loop as long as i is less than or equal to 100
i+=10 - Increase the loop counter value by 10 for each iteration

