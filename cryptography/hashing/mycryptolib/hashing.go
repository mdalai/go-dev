package mycryptolib

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func GetSha256Checksum(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		log.Fatal(err)
	}

	// get checksum in hex
	checksum := hash.Sum(nil)

	// convert hex to 64-digit string
	return hex.EncodeToString(checksum)
}

func GetMD5Checksum(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		log.Fatal(err)
	}

	// get checksum in hex 32-digit
	checksum := hash.Sum(nil)

	// convert hex to 32-digit string
	return hex.EncodeToString(checksum)
}
