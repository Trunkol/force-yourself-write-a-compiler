package main

import (
	"fmt"
	"force-yourself-write-a-compiler/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi, %s! Welcome to Stella REPL - Just for fun \n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
