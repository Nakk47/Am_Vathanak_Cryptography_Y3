package main

import "fmt"


func main() {
	var num1, num2 int

	// Read two integer inputs
	fmt.Print("Enter first integer: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second integer: ")
	fmt.Scan(&num2)

	fmt.Printf("\nEvaluating for num1 = %d, num2 = %d\n", num1, num2)
	fmt.Println("----------------------------------------")

	// 1. Both positive (&&)
	bothPositive := (num1 > 0) && (num2 > 0)
	fmt.Printf("Both positive          : %t\n", bothPositive)

	// 2. One greater than the other (||)
	oneGreater := (num1 > num2) || (num2 > num1)
	fmt.Printf("One greater than the other : %t\n", oneGreater)

	// 3. Not equal (!)
	notEqual := !(num1 == num2)
	fmt.Printf("Not equal                 : %t\n", notEqual)
}