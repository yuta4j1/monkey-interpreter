package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/yuta4j1/monkey-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Monkey programing language!\n", user.Username)

	fmt.Printf("Feel freee to type in command\n")
	repl.Start(os.Stdin, os.Stdout)
}
