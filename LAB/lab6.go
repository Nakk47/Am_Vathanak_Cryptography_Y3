package main

import "fmt"


// Function requested by the lab
func xorEncrypt(text string, key byte) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key // XOR each character byte with the key
	}
	return string(result)
}

func main() {
	var message string
	key := byte('K') // Single-byte key

	// 1. Take plaintext message input
	fmt.Print("Enter plaintext: ")
	fmt.Scan(&message)

	// 2. Encrypt
	encrypted := xorEncrypt(message, key)
	fmt.Printf("Key: %c (ASCII: %d)\n", key, key)
	fmt.Printf("Encrypted (Hex): %x\n", encrypted)

	// 3. Decrypt using the exact same function and key
	decrypted := xorEncrypt(encrypted, key)
	fmt.Printf("Decrypted: %s\n", decrypted)
}