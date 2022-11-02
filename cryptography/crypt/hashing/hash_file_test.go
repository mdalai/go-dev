package hashing

import (
	"testing"
	"os"
)


func TestGetMD5Checksum(t *testing.T) {
	testFile := "test.txt"
	strContent := "hello. this is test doc.\nplease calculate my checksum."
	_create_test_file(testFile, strContent)

	expected := "80ecf9401baeca471c34ccd4965c9fc2"
	got, _ := GetMD5Checksum(testFile)
	if expected != got {
		t.Fatalf(`GetMD5Hashing() test expected: %v, got: %v`, expected, got)
	}

	_rm_test_file(testFile)
}


func _create_test_file(fpath string, content string) {
	byteContent := []byte(content)
	if err := os.WriteFile(fpath, byteContent, 0644); err != nil {
		panic(err)
	}
}

func _rm_test_file(fpath string) {
	err := os.Remove(fpath)
	if err != nil {
		panic(err)
	}
}