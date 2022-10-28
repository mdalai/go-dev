package greetings

import (
	"fmt"
	"errors"
)

// Exported name is used for this function name
func Hello(name string) (string, error) {
	// error handle using pkg errors
	if name == "" {
		return "", errors.New("empty name")
	}
	// declare and initialze a var in one line
	// %v format verb
	msg := fmt.Sprintf("Hello, %v. Welcome!", name)
	return msg, nil
}