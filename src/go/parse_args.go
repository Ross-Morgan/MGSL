package main

import (
	"fmt"
	"os"
)

func main() {
	positionalArgs := [...]string{"1", "2", "3", "4"}

	args := os.Args[1:]

	fmt.Println(positionalArgs)
	fmt.Println(args)
}
