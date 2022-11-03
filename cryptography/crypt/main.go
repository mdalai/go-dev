package main

import (
	"fmt"
	"log"
	"flag"
	"os"

	"github.com/mdalai/go-dev/cryptography/crypt/encdec"
	"github.com/mdalai/go-dev/cryptography/crypt/hashing"
)

func parseArgs() {
	hashCmd := flag.NewFlagSet("hash", flag.ExitOnError)
	isSha256 := hashCmd.Bool("sha256", false, "use sha256")
	strVal := hashCmd.String("str", "", `"this string" need to be hashed`)
	encCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
	decCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Expected 'hash', 'encrypt' or 'decrypt' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "hash":
		hashCmd.Parse(os.Args[2:])
		var myHashVal string;
		var err error;
		if *strVal != "" {
			if *isSha256 {
				myHashVal = hashing.SHA256Hashing(*strVal)
			} else {
				myHashVal = hashing.MD5Hashing(*strVal)
			}
			fmt.Println(myHashVal)
		} else {
			for _, element := range hashCmd.Args() {
				if *isSha256 {
					myHashVal, err = hashing.GetSha256Checksum(element)
					if err != nil {
						log.Fatal(err)
					}
				} else {
					if myHashVal, err = hashing.GetMD5Checksum(element); err != nil {
						log.Fatal(err)
					}
				}		
				fmt.Println(myHashVal)
			}
		}
		
	case "encrypt":
		encCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'encrypt'")
		fmt.Println("  tail: ", encCmd.Args())	
	case "decrypt":
		hashCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'decrypt'")
		fmt.Println("  tail: ", decCmd.Args())
	default:
		fmt.Println("Expected 'hash', 'encrypt' or 'decrypt' subcommands")
		os.Exit(1)
	}

}

func main() {

	fmt.Println("-------Encrypt----------")

	encryptedByte := encdec.EncryptTxt([]byte("I have a secret need to be encrypted!"), "my.key.phrase")
	fmt.Println(encryptedByte)
	fmt.Println(string(encdec.EncryptTxt([]byte("My secret is that I don't have secret at all."), "my.key.phrase")))

	fmt.Println("-------Decrypt----------")
	decrypedByte := encdec.DecryptCipherTxt(encryptedByte, "my.key.phrase")
	fmt.Println(decrypedByte)
	fmt.Println(string(decrypedByte))

	fmt.Println("-------File checksum----------")
	parseArgs()
	
}
