# Go Slices
Slices are similar to arrays, but are more powerful and flexible.

Like arrays, slices are also used to store multiple values of the same type in a single variable.

However, unlike arrays, the length of a slice can grow and shrink as you see fit.

In Go, there are several ways to create a slice:

Using the []datatype{values} format
Create a slice from an array
Using the make() function

## Create a Slice With []datatype{values}
Syntax
slice_name := []datatype{values}
A common way of declaring a slice is like this:

myslice := []int{}
The code above declares an empty slice of 0 length and 0 capacity.

To initialize the slice during declaration, use this:

myslice := []int{1,2,3}
The code above declares a slice of integers of length 3 and also the capacity of 3.

In Go, there are two functions that can be used to return the length and capacity of a slice:

len() function - returns the length of the slice (the number of elements in the slice)
cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
Example
This example shows how to create slices using the []datatype{values} format:

package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}
Result:

0
0
[]
4
4
[Go Slices Are Powerful]
* In the example above, we see that in the first slice (myslice1), the actual elements are not specified, so both the length and capacity of the slice will be zero. In the second slice (myslice2), the elements are specified, and both length and capacity is equal to the number of actual elements specified.


## Create a Slice From an Array
You can create a slice by slicing an array:

Syntax
var myarray = [length]datatype{values} // An array
myslice := myarray[start:end] // A slice made from the array
* Example:
This example shows how to create a slice from an array:

package main
import ("fmt")

func main() {
  arr1 := [6]int{10, 11, 12, 13, 14,15}
  myslice := arr1[2:4]

  fmt.Printf("myslice = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice))
  fmt.Printf("capacity = %d\n", cap(myslice))
}
Result:

myslice = [12 13]
length = 2
capacity = 4

* In the example above myslice is a slice with length 2. It is made from arr1 which is an array with length 6.

The slice starts from the third element of the array which has value 12 (remember that array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.). The slice can grow to the end of the array. This means that the capacity of the slice is 4.
If myslice started from element 0, the slice capacity would be 6.

## Create a Slice With The make() Function
The make() function can also be used to create a slice.

Syntax
slice_name := make([]type, length, capacity)
Note: If the capacity parameter is not defined, it will be equal to length.

* Example:
```go
This example shows how to create slices using the make() function:

package main
import ("fmt")

func main() {
  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  // with omitted capacity
  myslice2 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))
}
Result:

myslice1 = [0 0 0 0 0]
length = 5
capacity = 10
myslice2 = [0 0 0 0 0]
length = 5
capacity = 5
```

