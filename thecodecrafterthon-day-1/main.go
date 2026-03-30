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
	for {
		fmt.Println("Enter first Number")
		_, error := fmt.Scanln(&a)
		if error != nil {
			fmt.Println("Invalid input: ")
			fmt.Println(error)
			continue
		}
		break
	}

	for {
		fmt.Println("Enter second number Number")
		_, error2 := fmt.Scanln(&b)
		if error2 != nil {
			fmt.Println("Invalid input: ")
			fmt.Println(error2)
			continue
		}
		break
	}

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
		if b != 0 {
			fmt.Println("your result is:", divide(a, b))
		}
		if b == 0 {
			fmt.Println("cannot divide by 0 :")
		}
		goto beginning
	default:
		fmt.Println("unknown operator")

	}

}
