package main

import (
	"log"
	"fmt"

	K "github.com/mdalai/go-dev/cryptography/keygen"
)

func main() {

	// Generate pub/pri key in pem file
	if err := K.GeneratePubPriKeyPairInPem("output/mykey"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n===== Keys in pem are generated!! ===== \n\n")

	
}