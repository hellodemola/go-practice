package main

import (
"fmt"
"errors"
 "math/rand"
)

func Hello (name string) (string, error) {
  if name == "" {
	return "", errors.New("Empty name")
  }
  message := fmt.Sprintf("Hi, %v. Welcome!", name)
  return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
