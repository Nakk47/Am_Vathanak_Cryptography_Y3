package main

import "fmt"

func main() {
	var a, b int

	// Read two integer inputs
	fmt.Print("Enter first number (a): ")
	fmt.Scan(&a)
	fmt.Print("Enter second number (b): ")
	fmt.Scan(&b)

	fmt.Printf("\nInitial values -> a: %d, b: %d\n", a, b)
	fmt.Println("----------------------------------------")

	// = (Simple Assignment)
	x := a
	fmt.Printf("1. Simple Assignment (=): x = %d\n", x)

	// += (Add and assign)
	x += b
	fmt.Printf("2. Addition Assignment (+= %d): x = %d\n", b, x)

	// -= (Subtract and assign)
	x -= b
	fmt.Printf("3. Subtraction Assignment (-= %d): x = %d\n", b, x)

	// *= (Multiply and assign)
	x *= b
	fmt.Printf("4. Multiplication Assignment (*= %d): x = %d\n", b, x)

	// /= (Divide and assign)
	if b != 0 {
		x /= b
		fmt.Printf("5. Division Assignment (/= %d): x = %d\n", b, x)
	} else {
		fmt.Println("5. Division Assignment: Skipped (cannot divide by 0)")
	}

	// %= (Modulus and assign)
	if b != 0 {
		x %= b
		fmt.Printf("6. Modulus Assignment (%%= %d): x = %d\n", b, x)
	} else {
		fmt.Println("6. Modulus Assignment: Skipped (cannot divide by 0)")
	}
}