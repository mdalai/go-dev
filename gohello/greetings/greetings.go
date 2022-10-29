package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// Exported name is used for this function name
func Hello(name string) (string, error) {
	// error handle using pkg errors
	if name == "" {
		return "", errors.New("empty name")
	}
	// declare and initialze a var in one line
	// %v format verb
	//msg := fmt.Sprintf("Hello, %v. Welcome!", name)
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

// Can Hello to multiple names
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	// why use make() for map, but not for slice?
	msgs := make(map[string]string)

	// How to loop slice?
	for _, name := range names {
		// re-use previous func
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		msgs[name] = msg
	}

	return msgs, nil
}


// init sets initial values for variables used in the function
func init() {
	// why seed?
	rand.Seed(time.Now().UnixNano())
}


// return message is selected at random
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hello, %v. Welcome!",
		"Great to have you, %v!",
		"Sain uu, %v!",
		"Hail, %v Well met!",
	}
	// randomly select a item of formats per random index within [0, 4]
	return formats[rand.Intn(len(formats))]
}