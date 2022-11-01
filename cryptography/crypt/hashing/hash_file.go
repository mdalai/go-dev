package hashing

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"errors"
	"os"
	"fmt"
)

func GetSha256Checksum(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		errMsg := fmt.Sprintf("File open error: %v", filepath)
		return "", errors.New(errMsg)
	}
	defer f.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	// get checksum in hex
	checksum := hash.Sum(nil)

	// convert hex to 64-digit string
	return hex.EncodeToString(checksum), nil
}

func GetMD5Checksum(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	// get checksum in hex 32-digit
	checksum := hash.Sum(nil)

	// convert hex to 32-digit string
	return hex.EncodeToString(checksum), nil
}
