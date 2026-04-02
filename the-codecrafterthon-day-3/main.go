// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Your Name: Amina Oteikwu(aoteikwu).]
// Squad:  [Your Squad Name: The Pointers.]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var smallWords = map[string]bool{
	"a": true, "an": true, "the": true,
	"and": true, "but": true, "or": true,
	"for": true, "nor": true, "on": true,
	"at": true, "to": true, "by": true,
	"in": true, "of": true, "up": true,
	"as": true, "is": true, "it": true,
}

func main() {
	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE WELCOM ON BOARD, LET's INTERACT.")
	fmt.Println("──────────────────────────────────────")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := strings.ToLower(parts[0])

		if command == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
		}

		if len(parts) < 2 {
			fmt.Printf("✗ No text provided. Usage: %s <text>\n", command)
			continue
		}

		text := strings.Join(parts[1:], " ")
		var result string

		switch command {
		case "upper":
			result = ToUpper(text)
		case "lower":
			result = ToLower(text)
		case "cap":
			result = ToCap(text)
		case "title":
			result = ToTitle(text)
		case "snake":
			result = ToSnake(text)
		case "reverse":
			result = ReverseWords(text)
		default:
			fmt.Printf("✗ Unknown command: \"%s\"\n", command)
			fmt.Println("  Valid commands: upper, lower, cap, title, snake, reverse, exit")
			continue
		}

		fmt.Printf("→ %s\n", result)
	}
}

func ToUpper(text string) string {
	return strings.ToUpper(text)
}

func ToLower(text string) string {
	return strings.ToLower(text)
}

func ToCap(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}

func ToTitle(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		lw := strings.ToLower(w)
		if i == 0 || !smallWords[lw] {
			words[i] = strings.ToUpper(string(lw[0])) + lw[1:]
		} else {
			words[i] = lw
		}
	}
	return strings.Join(words, " ")
}

func ToSnake(text string) string {
	var sb strings.Builder
	text = strings.ToLower(text)
	for _, ch := range text {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			sb.WriteRune(ch)
		} else if ch == ' ' || ch == '-' {
			sb.WriteRune('_')
		}

	}
	return sb.String()
}

func ReverseWords(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		runes := []rune(w)
		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}
