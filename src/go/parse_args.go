package main

import (
	"fmt"
	"os"
)

func argShorthand(argName string) string {
	return fmt.Sprintf("-%s", string(argName))
}

func checkEnoughArgs(required []string, given []string) bool {
	return len(given) >= len(required)
}

func exitWithCode(code int) {
	fmt.Fprint(os.Stderr, "Exiting with error code ")
	fmt.Fprint(os.Stderr, code)
	fmt.Fprint(os.Stderr, " (")
	fmt.Fprintf(os.Stderr, "0x%x", code)
	fmt.Fprint(os.Stderr, ")")
	fmt.Fprintln(os.Stderr, "")
	os.Exit(code)
}

func main() {
	// argMap := map[string]string{}

	requiredArgs := []string{"filename"}
	positionalArgs := []string{"verbose"}

	args := os.Args[1:]

	enough := checkEnoughArgs(requiredArgs, args)

	if !(enough) {
		fmt.Fprintln(os.Stderr, "Not enough args:")

		fmt.Fprint(os.Stderr, "Required: ")
		fmt.Fprintln(os.Stderr, len(requiredArgs))

		fmt.Fprint(os.Stderr, "Provided: ")
		fmt.Fprintln(os.Stderr, len(args))

		exitWithCode(1)
	}

	fmt.Println(requiredArgs)
	fmt.Println(positionalArgs)
	fmt.Println(args)
}
