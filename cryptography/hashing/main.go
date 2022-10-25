package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"

	"fmt"
	"io"
	"log"
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


func EncryptTxt(byteTxt []byte , key string) []byte {
	// create an AES block
	aesBlock, err := aes.NewCipher([]byte(MD5Hashing(key)) )
	if err != nil {
		fmt.Println(err)
	}

	// create a new 128bit cipher with a nonce using an AES block
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	// create cipher text
	cipherText := gcmInstance.Seal(nonce, nonce, byteTxt, nil)

	return cipherText
}


func DecryptCipherTxt(cipherTxt []byte, key string) []byte {
	// hash the key
	hashedKey := MD5Hashing(key)
	// create a AES block
	aesBlock, err := aes.NewCipher([]byte(hashedKey))
	if err != nil {
		log.Fatalln(err)
	}
	// create a new cipher instance
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		log.Fatalln(err)
	}

	// split nonce and cipher text
	nonceSize := gcmInstance.NonceSize()
	nonce, cipherText := cipherTxt[:nonceSize], cipherTxt[nonceSize:]

	// decrypt cipher text
	origText, err := gcmInstance.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return origText
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


	fmt.Println("-------Encrypt----------")

	encryptedByte := EncryptTxt([]byte("I have a secret need to be encrypted!"), "my.key.phrase")
	fmt.Println(encryptedByte)
	fmt.Println(string(EncryptTxt([]byte("My secret is that I don't have secret at all."), "my.key.phrase")))

	fmt.Println("-------Decrypt----------")
	decrypedByte := DecryptCipherTxt(encryptedByte, "my.key.phrase")
	fmt.Println(decrypedByte)
	fmt.Println(string(decrypedByte))
	
}