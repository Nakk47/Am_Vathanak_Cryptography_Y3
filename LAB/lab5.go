package main

import "encoding/base64"
import "fmt"

func main() {
	var text string

	// 1. Take string input from the user
	fmt.Print("Enter text: ")
	fmt.Scan(&text)

	// 2. Display Binary representation
	fmt.Print("Binary: ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%08b ", text[i])
	}
	fmt.Println()

	// 3. Display Hexadecimal representation
	fmt.Printf("Hexadecimal: %x\n", text)

	// 4. Display Base64 representation
	base64 := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Printf("Base64: %s\n", base64)
}