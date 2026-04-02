package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertNumber(value string, base string) {
	switch base {
	case "hex":

		if _, err := strconv.ParseInt(value, 16, 64); err != nil {
			fmt.Println("Invalid hex input.")
			return
		}
		decimalValue, _ := strconv.ParseInt(value, 16, 64)
		fmt.Printf("✦ Decimal: %d\n", decimalValue)

	case "bin":

		if _, err := strconv.ParseInt(value, 2, 64); err != nil {
			fmt.Println("Invalid binary input.")
			return
		}
		decimalValue, _ := strconv.ParseInt(value, 2, 64)
		fmt.Printf("✦ Decimal: %d\n", decimalValue)

	case "dec":

		decimalValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			fmt.Println("Invalid decimal input.")
			return
		}
		binaryValue := strconv.FormatInt(decimalValue, 2)
		hexValue := strconv.FormatInt(decimalValue, 16)
		fmt.Printf("✦ Binary:  %s\n", binaryValue)
		fmt.Printf("✦ Hex:     %s\n", strings.ToUpper(hexValue))

	default:
		fmt.Println("Base must be 'hex', 'bin', or 'dec'.")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter value and base (e.g., '1E hex') or 'quit' to exit: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "quit" {
			fmt.Println("Exiting the program.")
			break
		}

		parts := strings.Fields(input)

		if len(parts) != 2 {
			fmt.Println("Please provide a value and a base.")
			continue
		}

		value := parts[0]
		base := strings.ToLower(parts[1])

		convertNumber(value, base)
	}
}
