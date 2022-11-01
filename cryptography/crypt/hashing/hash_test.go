package hashing

import (
	"testing"

	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func TestMD5HashingLength(t *testing.T) {
	tstStr := "Test me"
	expected := 32
	tstStrHash := MD5Hashing(tstStr)
	if expected != len(tstStrHash) {
		t.Fatalf(`MD5Hashing() length expected: %v, got: %v`, expected, len(tstStrHash))
	}	
}

func TestMD5Hashing(t *testing.T) {
	tstStr := "What is my hash"
	tstStrHashedByte := md5.Sum([]byte(tstStr))
	expected := hex.EncodeToString(tstStrHashedByte[:])
	got := MD5Hashing(tstStr)
	if expected != got {
		t.Fatalf(`MD5Hashing() test expected: %v, got: %v`, expected, got)
	}
}

func TestSHA256HashingLength(t *testing.T) {
	tstStr := "Test me"
	expected := 64
	tstStrHash := SHA256Hashing(tstStr)
	if expected != len(tstStrHash) {
		t.Fatalf(`SHA256Hashing() length expected: %v, got: %v`, expected, len(tstStrHash))
	}	
}

func TestSHA256Hashing(t *testing.T) {
	tstStr := "What is my hash"
	tstStrHashedByte := sha256.Sum256([]byte(tstStr))
	expected := hex.EncodeToString(tstStrHashedByte[:])
	got := SHA256Hashing(tstStr)
	if expected != got {
		t.Fatalf(`SHA256Hashing() test expected: %v, got: %v`, expected, got)
	}
}