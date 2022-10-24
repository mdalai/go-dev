package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"

	"fmt"
)

// Hasing algorithm is one-way function. It is used for validation.
// For communication, use encryption algorithm like AES.

// The Message Digest Method 5 Algorithm (MD5) is 
// a cryptographic algorithm that returns a 128-bit digest 
// from any text represented as 32-digit hexadecimal.
func MD5Hashing(inputStr string) string {
	// string to bytes
	byteInput := []byte(inputStr)
	// md5 hash returns 32-digit hexadecimal
	md5Hash := md5.Sum(byteInput)
	// hexadecimal to string
	return hex.EncodeToString(md5Hash[:])
}

// The SHA256 (Secure Hash Algorithm 256-bits) is 
// a cryptographic one-way hashing algorithm that returns a 256-bit hash value. 
func SHA256Hashing(inputStr string) string {
	// string to byte
	byteInput := []byte(inputStr)
	// returns 64-digit hex
	sha256Hash := sha256.Sum256(byteInput)
	// hex to string
	return hex.EncodeToString(sha256Hash[:])
}


func main()  {
	myStr := "hello world"

	fmt.Println(myStr)
	fmt.Println(MD5Hashing(myStr))
	fmt.Println(SHA256Hashing(myStr))

	myStr = "Hashing strings using GO crypto package"

	fmt.Println(myStr)
	fmt.Println(MD5Hashing(myStr))
	fmt.Println(SHA256Hashing(myStr))
	
}