package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}
func sub(a, b int) int {
	return a - b
}
func divide(a, b int) int {
	return a / b
}

func main() {
beginning:
	var choice string
	fmt.Println("choose an operation: 1. add 2. mul 3.sub 4. divide 5. quit")
	fmt.Scanln(&choice)
	var a int
	var b int
	if choice == "5" {
		fmt.Println("bye bye!")
		return
	}
	fmt.Println("Enter a Number")
	fmt.Scanln(&a)
	fmt.Println("Enter Second Number")
	fmt.Scanln(&b)

	switch choice {
	case "1":
		fmt.Println("your result is:", add(a, b))

		goto beginning

	case "2":

		fmt.Println("your result is:", sub(a, b))
		goto beginning

	case "3":

		fmt.Println("your result is:", mul(a, b))
		goto beginning

	case "4":

		fmt.Println("your result is:", divide(a, b))
		goto beginning

	}
}
