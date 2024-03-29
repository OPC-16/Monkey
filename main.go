package main

import (
    "os"
    "os/user"
    "fmt"

    "github.com/OPC-16/Monkey/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s! This is Monkey Programming Language!\n", user.Username)
    fmt.Printf("Feel free to type in commands\n")
    repl.Start(os.Stdin, os.Stdout)
}
