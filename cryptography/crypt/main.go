package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"strings"
	"strconv"

	"github.com/mdalai/go-dev/cryptography/crypt/encdec"
	"github.com/mdalai/go-dev/cryptography/crypt/hashing"
)

func parseArgs() {
	hashCmd := flag.NewFlagSet("hash", flag.ExitOnError)
	isSha256 := hashCmd.Bool("sha256", false, "Optional. If hash with sha256, use this boolean argument. Otherwise hashes with MD5 by default.")
	strVal := hashCmd.String("str", "", `Optional. If hash a string, provide "the string" using this argument.`)
	
	encCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
	enc2bytes := encCmd.Bool("toByte", false, "Optional. If prefer encrypting to byte, use this argument. Otherwise, encrypted string is returned.")
	encStr := encCmd.String("str", "", "Optional. If encrypt a string, provide it using this argument.")

	decCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	dec2bytes := decCmd.Bool("toByte", false, "Optional. If prefer decrypting to byte, use this argument. Otherwise, decrypted string is returned.")
	decStr := decCmd.String("str", "", "Optional. If decrypt a string, provide it using this argument.")


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
		var encryptedByte []byte
		if *encStr != "" {
			fmt.Println("Encrypting this string: ", *encStr)
			encryptedByte = encdec.EncryptTxt([]byte(*encStr), "my.key.phrase")
		}
		if *enc2bytes {
			fmt.Println(encryptedByte)
		} else {
			fmt.Println(string(encryptedByte))
		}
	case "decrypt":
		decCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'decrypt'")
		fmt.Println("  tail: ", decCmd.Args())
		var decryptedByte []byte
		if *decStr != "" {
			fmt.Println("Decrypting this string: ", *decStr)
			fmt.Println("Converted to []byte: ", byteStr2bytes(*decStr))
			decryptedByte = encdec.DecryptCipherTxt(byteStr2bytes(*decStr), "my.key.phrase")
		}
		if *dec2bytes {
			fmt.Println(decryptedByte)
		} else {
			fmt.Println(string(decryptedByte))
		}
	default:
		fmt.Println("Expected 'hash', 'encrypt' or 'decrypt' subcommands")
		os.Exit(1)
	}

}

func byteStr2bytes(bStr string) []byte {
	var bb []byte
	for _, byteNstr := range strings.Split(strings.Trim(bStr, "[]"), " ") {
		byteN,_ := strconv.Atoi(byteNstr)
		bb = append(bb,byte(byteN))
	}
	return bb
}

func main() {
	parseArgs()
}
