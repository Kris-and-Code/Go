package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Operation represents a calculator operation
type Operation struct {
	name     string
	symbol   string
	function func(float64, float64) (float64, error)
}

// Calculator holds the available operations
type Calculator struct {
	operations map[int]Operation
}

// NewCalculator creates a new calculator instance
func NewCalculator() *Calculator {
	return &Calculator{
		operations: map[int]Operation{
			1: {
				name:   "Addition",
				symbol: "+",
				function: func(a, b float64) (float64, error) {
					return a + b, nil
				},
			},
			2: {
				name:   "Subtraction",
				symbol: "-",
				function: func(a, b float64) (float64, error) {
					return a - b, nil
				},
			},
			3: {
				name:   "Multiplication",
				symbol: "*",
				function: func(a, b float64) (float64, error) {
					return a * b, nil
				},
			},
			4: {
				name:   "Division",
				symbol: "/",
				function: func(a, b float64) (float64, error) {
					if b == 0 {
						return 0, fmt.Errorf("division by zero")
					}
					return a / b, nil
				},
			},
		},
	}
}

// readInput reads a line of input from the user
func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// getNumber prompts for and reads a number from the user
func getNumber(prompt string) (float64, error) {
	input := readInput(prompt)
	return strconv.ParseFloat(input, 64)
}

// displayMenu shows the calculator menu
func (c *Calculator) displayMenu() {
	fmt.Println("\nSelect operation:")
	for i := 1; i <= len(c.operations); i++ {
		fmt.Printf("%d. %s\n", i, c.operations[i].name)
	}
	fmt.Println("5. Exit")
}

// run performs the calculator operations
func (c *Calculator) run() {
	fmt.Println("Simple Calculator")
	fmt.Println("----------------")

	for {
		c.displayMenu()
		choiceStr := readInput("Enter your choice (1-5): ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue
		}

		if choice == 5 {
			fmt.Println("Thank you for using the calculator!")
			return
		}

		operation, exists := c.operations[choice]
		if !exists {
			fmt.Println("Invalid choice! Please try again.")
			continue
		}

		num1, err := getNumber("Enter first number: ")
		if err != nil {
			fmt.Println("Invalid number! Please try again.")
			continue
		}

		num2, err := getNumber("Enter second number: ")
		if err != nil {
			fmt.Println("Invalid number! Please try again.")
			continue
		}

		result, err := operation.function(num1, num2)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("\nResult: %.2f %s %.2f = %.2f\n", num1, operation.symbol, num2, result)
	}
}

func main() {
	calculator := NewCalculator()
	calculator.run()
} 