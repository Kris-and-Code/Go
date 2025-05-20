package main

import (
	"bufio"
	"fmt"
	"math"
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

// Calculation represents a single calculation entry
type Calculation struct {
	num1    float64
	num2    float64
	op      string
	result  float64
}

// Calculator holds the available operations and state
type Calculator struct {
	operations map[int]Operation
	memory     float64
	history    []Calculation
	precision  int
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
			5: {
				name:   "Power",
				symbol: "^",
				function: func(a, b float64) (float64, error) {
					return math.Pow(a, b), nil
				},
			},
			6: {
				name:   "Square Root",
				symbol: "√",
				function: func(a, b float64) (float64, error) {
					if a < 0 {
						return 0, fmt.Errorf("cannot calculate square root of negative number")
					}
					return math.Sqrt(a), nil
				},
			},
			7: {
				name:   "Percentage",
				symbol: "%",
				function: func(a, b float64) (float64, error) {
					return (a * b) / 100, nil
				},
			},
		},
		memory:    0,
		history:   make([]Calculation, 0),
		precision: 2,
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
	fmt.Println("8. Memory Functions")
	fmt.Println("9. View History")
	fmt.Println("10. Set Precision")
	fmt.Println("11. Exit")
}

// displayMemoryMenu shows the memory functions menu
func (c *Calculator) displayMemoryMenu() {
	fmt.Println("\nMemory Functions:")
	fmt.Println("1. Store in Memory (M+)")
	fmt.Println("2. Recall from Memory (MR)")
	fmt.Println("3. Clear Memory (MC)")
	fmt.Println("4. Back to Main Menu")
}

// handleMemoryFunctions processes memory-related operations
func (c *Calculator) handleMemoryFunctions() {
	for {
		c.displayMemoryMenu()
		choiceStr := readInput("Enter your choice (1-4): ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			num, err := getNumber("Enter number to store: ")
			if err != nil {
				fmt.Println("Invalid number! Please try again.")
				continue
			}
			c.memory = num
			fmt.Printf("Stored %.2f in memory\n", num)
		case 2:
			fmt.Printf("Memory value: %.2f\n", c.memory)
		case 3:
			c.memory = 0
			fmt.Println("Memory cleared")
		case 4:
			return
		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}

// setPrecision allows the user to set the decimal precision
func (c *Calculator) setPrecision() {
	precisionStr := readInput("Enter decimal precision (0-10): ")
	precision, err := strconv.Atoi(precisionStr)
	if err != nil || precision < 0 || precision > 10 {
		fmt.Println("Invalid precision! Please enter a number between 0 and 10.")
		return
	}
	c.precision = precision
	fmt.Printf("Precision set to %d decimal places\n", precision)
}

// displayHistory shows the calculation history
func (c *Calculator) displayHistory() {
	if len(c.history) == 0 {
		fmt.Println("No calculations in history")
		return
	}
	fmt.Println("\nCalculation History:")
	for i, calc := range c.history {
		fmt.Printf("%d. %.2f %s %.2f = %.2f\n", i+1, calc.num1, calc.op, calc.num2, calc.result)
	}
}

// run performs the calculator operations
func (c *Calculator) run() {
	fmt.Println("Advanced Calculator")
	fmt.Println("------------------")

	for {
		c.displayMenu()
		choiceStr := readInput("Enter your choice (1-11): ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue
		}

		switch choice {
		case 11:
			fmt.Println("Thank you for using the calculator!")
			return
		case 8:
			c.handleMemoryFunctions()
			continue
		case 9:
			c.displayHistory()
			continue
		case 10:
			c.setPrecision()
			continue
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

		var num2 float64
		if choice != 6 { // Square root only needs one number
			num2, err = getNumber("Enter second number: ")
			if err != nil {
				fmt.Println("Invalid number! Please try again.")
				continue
			}
		}

		result, err := operation.function(num1, num2)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// Store calculation in history
		c.history = append(c.history, Calculation{
			num1:   num1,
			num2:   num2,
			op:     operation.symbol,
			result: result,
		})

		// Format result with specified precision
		format := fmt.Sprintf("%%.%df", c.precision)
		fmt.Printf("\nResult: %s %s %s = %s\n",
			fmt.Sprintf(format, num1),
			operation.symbol,
			fmt.Sprintf(format, num2),
			fmt.Sprintf(format, result))
	}
}

func main() {
	calculator := NewCalculator()
	calculator.run()
} 