package main

import "fmt"

func myAND(a, b int) {
	fmt.Println("AND (a & b):", a&b)
}

func myOR(a, b int) {
	fmt.Println("OR (a | b):", a|b)
}

func myXOR(a, b int) {
	fmt.Println("XOR (a ^ b):", a^b)
}

func myNOT(a, b int) {
	// In Go, bitwise NOT is the ^ operator (unary)
	fmt.Println("NOT (^a):", ^a)
	fmt.Println("NOT (^b):", ^b)
}

func myShifts(a int) {
	fmt.Println("Left Shift (a << 1):", a<<1)
	fmt.Println("Right Shift (a >> 1):", a>>1)
}

func main() {
	a := 12 // binary: 1100
	b := 5  // binary: 0101

	fmt.Println("Original values: a =", a, "b =", b)
	fmt.Println("--- Bitwise Functions ---")

	// 1. Bitwise functions from pseudocode
	myXOR(a, b)
	myNOT(a, b)
	myOR(a, b)
	myAND(a, b)
	myShifts(a)

	fmt.Println("--- Assignment Operators ---")

	// 2. Effect of bitwise assignment operators on variable values
	x := a
	x &= b // x = x & b
	fmt.Println("After x &= b:", x)

	x = a
	x |= b // x = x | b
	fmt.Println("After x |= b:", x)

	x = a
	x ^= b // x = x ^ b
	fmt.Println("After x ^= b:", x)
}