package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Snehil-Shah/target/internal/parser"
	"github.com/Snehil-Shah/target/internal/validator"
)

func main() {
	fmt.Println("Go Fuzzing Sample")
	fmt.Println("Enter commands in format: [action] [input]")
	fmt.Println("Actions: parse, parsejson, email, numeric, exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)

		if len(parts) < 1 {
			continue
		}

		cmd := parts[0]
		input := ""
		if len(parts) > 1 {
			input = parts[1]
		}

		switch cmd {
		case "exit", "quit":
			return

		case "parse":
			result, err := parser.ParseInput(input)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println(result)
			}

		case "parsejson":
			result, err := parser.ParseJSON(input)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Parsed JSON:")
				for k, v := range result {
					fmt.Printf("  %s: %s\n", k, v)
				}
			}

		case "email":
			err := validator.ValidateEmail(input)
			if err != nil {
				fmt.Printf("Invalid email: %v\n", err)
			} else {
				fmt.Println("Valid email!")
			}

		case "numeric":
			value, unit, err := validator.ParseNumeric(input)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Value: %f, Unit: %s\n", value, unit)
			}

		default:
			fmt.Println("Unknown command")
		}
	}
}
