package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi Prakyath. I'm a calculator, I can help you with any calculations you need.")

	for {
		// Read input from user
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter calculation: ")
		text, _ := reader.ReadString('\n')

		// trim the newline character from the input
		text = strings.TrimSpace(text)

		// check if the user entered "exit" to exit the program
		if text == "exit" {
			break
		} 

		// split the input into 2 parts - right and left operand 
		parts := strings.Split(text, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input. Please enter like '2 + 3'")
			continue
		}

		// convert the operands to integers
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid input. Try again")
			continue
		}
		right, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Invalid input. Try again")
			continue
		}

		// Perform calculation based on operator
		var result int
		switch parts[1] {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right		
		case "/":
			result = left / right
		default:
			fmt.Println("Invalid operator. Please use +, -, *, or /")
			continue
		}

		// Print the result
		fmt.Printf("Result: %d\n", result)

	}
}