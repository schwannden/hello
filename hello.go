package main

import (
    "fmt"
    "log"

    "github.com/schwannden/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Get a greeting message and print it.
    message, err := greetings.Hello("Schwannden")
    // if error occur, print error message and exit
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)

    names := []string{"Gladys", "Samantha", "Darrin"}
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(messages)
}
