package main

import (
	"github.com/pippokairos/pizza/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
