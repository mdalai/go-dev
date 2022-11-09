package encdec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"

	"github.com/mdalai/go-dev/cryptography/crypt/hashing"
)

// Block cipher algorithms are AES (Advanced Encryption System), DES (Data Encryption Standart), RC6 (Rives Cipher 6) ...

func EncryptTxt(plainTxt []byte, key string) []byte {

	// Hash the secret key as a more layer of protection
	hashedKey := []byte(hashing.MD5Hashing(key))

	// create an AES cipher block using the hashed secret key
	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		log.Fatalf("AES cipher err: %v", err.Error())
	}

	// The GCM (Galois/Counter Mode) is a stream mode and provides data authenticity and confidentially.
	// create a new 128bit cipher with a nonce using an AES block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("AES cipher GCM err: %v", err.Error())
	}

	// Generate random Nonce which will be used for Seal function
	// The nonce has to be unique and it changes every time when data is encrypted.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
	}

	// Encrypt data using Seal function
	cipherText := gcm.Seal(nonce, nonce, plainTxt, nil)

	return cipherText
}

func DecryptCipherTxt(cipherTxt []byte, key string) []byte {
	// hash the key
	hashedKey := hashing.MD5Hashing(key)
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
