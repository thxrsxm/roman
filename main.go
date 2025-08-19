package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thxrsxm/roman/internal/roman"
)

// main runs the Roman numeral converter CLI. It processes a single command-line
// argument as input, if provided, or enters an interactive mode reading from stdin
// until "exit" is entered.
func main() {
	args := os.Args
	if len(args) > 1 {
		printRoman(args[1])
		os.Exit(0)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("roman: ")
		// Read until newline
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			fmt.Println()
			continue
		}
		// Clean up input (remove newline and surrounding spaces)
		userInput = strings.TrimSpace(userInput)
		switch userInput {
		case "exit":
			os.Exit(0)
		case "info":
			fmt.Println("Created by Erik Andrè Thürsam.")
		default:
			printRoman(userInput)
		}
		fmt.Println()
	}
}

// printRoman converts the input string to an integer and prints its Roman numeral
// representation. It prints an error message if the input is not a valid integer.
func printRoman(input string) {
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println(roman.ConvertToRoman(num))
}
