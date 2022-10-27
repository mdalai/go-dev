package encdec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"

	"github.com/mdalai/go-dev/cryptography/crypt/hashing"
)

func EncryptTxt(byteTxt []byte, key string) []byte {
	// create an AES block
	aesBlock, err := aes.NewCipher([]byte(hashing.MD5Hashing(key)))
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
