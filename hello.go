package main

import (
"fmt"
"log"
)


func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
    names := []string{"Gladys", "Samantha", "Darrin"}
    message, err := Hellos(names)

	if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
    mains()
}
