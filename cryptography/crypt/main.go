package main

import (
	"fmt"
	"log"

	"github.com/mdalai/go-dev/cryptography/crypt/encdec"
	"github.com/mdalai/go-dev/cryptography/crypt/hashing"
)

func main() {
	myStr := "hello world"

	fmt.Println(myStr)
	fmt.Println(hashing.MD5Hashing(myStr))
	fmt.Println(hashing.SHA256Hashing(myStr))

	myStr = "Hashing strings using GO crypto package"

	fmt.Println(myStr)
	fmt.Println(hashing.MD5Hashing(myStr))
	fmt.Println(hashing.SHA256Hashing(myStr))

	fmt.Println("-------Encrypt----------")

	encryptedByte := encdec.EncryptTxt([]byte("I have a secret need to be encrypted!"), "my.key.phrase")
	fmt.Println(encryptedByte)
	fmt.Println(string(encdec.EncryptTxt([]byte("My secret is that I don't have secret at all."), "my.key.phrase")))

	fmt.Println("-------Decrypt----------")
	decrypedByte := encdec.DecryptCipherTxt(encryptedByte, "my.key.phrase")
	fmt.Println(decrypedByte)
	fmt.Println(string(decrypedByte))

	fmt.Println("-------File checksum----------")
	myHashVal, err := hashing.GetSha256Checksum("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myHashVal)

	if myHashVal, err = hashing.GetMD5Checksum("test.txt"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(myHashVal)
}
