package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Simple Calculator")
	fmt.Println("----------------")

	for {
		fmt.Println("\nSelect operation:")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice (1-5): ")
		fmt.Scan(&choice)

		if choice == 5 {
			fmt.Println("Thank you for using the calculator!")
			os.Exit(0)
		}

		if choice < 1 || choice > 4 {
			fmt.Println("Invalid choice! Please try again.")
			continue
		}

		var num1, num2 float64
		fmt.Print("Enter first number: ")
		fmt.Scan(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)

		var result float64
		var operation string

		switch choice {
		case 1:
			result = num1 + num2
			operation = "+"
		case 2:
			result = num1 - num2
			operation = "-"
		case 3:
			result = num1 * num2
			operation = "*"
		case 4:
			if num2 == 0 {
				fmt.Println("Error: Cannot divide by zero!")
				continue
			}
			result = num1 / num2
			operation = "/"
		}

		fmt.Printf("\nResult: %.2f %s %.2f = %.2f\n", num1, operation, num2, result)
	}
} 