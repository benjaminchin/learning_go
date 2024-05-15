package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the Logger,
    // including the log entry prefix and a flag to disable
    // printing the time, source file and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)


    // A slice of names.
    names := []string{"Ben", "Mae", "Owen"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)

    // Get a greeting and print it
    // message, err := greetings.Hello("Ben")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(messages)
}
