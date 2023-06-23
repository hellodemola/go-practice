package main

import (
"fmt"
"log"
)


func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
    message, err := Hello("Ademola")

	if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}
