package main

import (
	"fmt"
)

// Functions for arithmetic operations
func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) {
	if b == 0 {
		fmt.Println("Error: division by zero")
		return
	}
	fmt.Printf("Result: %d\n", a/b)
}

func mod(a, b int) {
	if b == 0 {
		fmt.Println("Error: division by zero")
		return
	}
	fmt.Printf("Result: %d\n", a%b)
}

func main() {
	var choice int
	var a, b int

	// Infinite loop to repeat until the user selects Exit
	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("  1) Add  2) Sub  3) Mul  4) Div  5) Mod  6) Exit")
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		// Check for Exit first
		if choice == 6 {
			fmt.Println("Done")
			break
		}

		// Handle invalid menu selection
		if choice < 1 || choice > 6 {
			fmt.Println("Error: invalid option, please choose between 1 and 6.")
			continue
		}

		// Read the two numbers from the user
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		// Execute the chosen operation
		switch choice {
		case 1:
			fmt.Printf("Result: %d\n", add(a, b))
		case 2:
			fmt.Printf("Result: %d\n", sub(a, b))
		case 3:
			fmt.Printf("Result: %d\n", mul(a, b))
		case 4:
			div(a, b)
		case 5:
			mod(a, b)
		}
	}
}