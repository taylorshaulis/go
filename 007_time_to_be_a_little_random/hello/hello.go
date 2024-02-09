package main

import (
    "fmt"
    "log"
    // "math/rand"
    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("beautiful")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)


    // formats := []string{
    //     "Hi, %v. Welcome!",
    //     "Great to see you, %v!",
    //     "Hail, %v! Well met!",
    // }

    // fmt.Println(len(formats))
    // fmt.Println(rand.Intn(3))
}