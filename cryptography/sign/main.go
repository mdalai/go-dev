package main

import (
	"fmt"
	"log"

	S "github.com/mdalai/go-dev/cryptography/sign/signing"
	K "github.com/mdalai/go-dev/cryptography/keygen"
)

func main() {
	fmt.Println("Start signing...")

	// your text here
	plainTxt := []byte("I am who i am. Sign me")

	// Create RSA key pair: priKey, pubKey
	priKey, pubKey, err := K.GeneratePubPriKeyPair()
	if err != nil {
		log.Fatal(err)
	}
	
	//fmt.Printf("\n===== privateKey ===== \n\n %v \n\n", priKey)
	fmt.Printf("\n===== publicKey ===== \n\n %v \n\n", pubKey)

	// Create a signature
	signature, err := S.SignTxt(plainTxt, priKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n===== signature ===== \n\n %v \n\n", signature)


	// Verify signature
	err = S.VerifyTxt(plainTxt, pubKey, signature)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Signature verification success!")

}