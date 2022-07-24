package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Colors representing different program statuses
var (
	Info    = Teal
	Warn    = Yellow
	Success = Green
	Fatal   = Red
)

// ANSI color codes
var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

// Convert ANSI formatting sequences to string-wrapping functions
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString, fmt.Sprint(args...))
	}
	return sprint
}

// Apply a function to every element of an array
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Remove all duplicate elements from an array
func RemoveDuplicates(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

// Print information about mission programs
func LogMissing(missingCommands []string) int {
	if len(missingCommands) == 0 {
		fmt.Println("No programs missing!")
		return 0
	}

	fmt.Println("\nMissing Programs:")

	for _, command := range missingCommands {
		fmt.Println(command)
	}

	fmt.Println("\nInstall / Modify these programs so that they can be called in the shell.")

	return 1
}

// Return first word of a string
func FirstWord(s string) string {
	return strings.Fields(s)[0]
}

// Test if program is accessible from the shell
func CommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// Test if several programs are accesible from the shell
func CheckCommandsExist(commands []string) int {
	fmt.Println("Checking program existence")
	fmt.Println("----------------------------")

	missing := []string{}

	commandWords := Map(commands, FirstWord)
	commandWords = RemoveDuplicates(commandWords)

	var commandExistence bool
	var printColor func(...interface{}) string

	for _, command := range commandWords {
		commandExistence = CommandExists(command)

		if commandExistence {
			printColor = Success
		} else {
			missing = append(missing, command)
			printColor = Fatal
		}

		fmt.Printf("%s\t|\t\t%s\n", command, printColor(strconv.FormatBool(commandExistence)))
	}

	fmt.Println("----------------------------")

	exitCode := LogMissing(missing)

	fmt.Println("----------------------------")

	return exitCode
}

// Run commands in shell
func RunCommands(commands []string) {
	for _, command := range commands {
		err := exec.Command(command).Run()

		if err != nil {
			panic(err)
		}
	}
}

func exitWithCode(code int) {
	fmt.Fprintf(os.Stderr, "(0x%x) ", code)
	os.Exit(code)
}

// Program entry point
func main() {
	buildCommands := []string{
		"g++ c++/main.cpp -o ../bin/mgsl.exe",
		"g++ c++/shell.cpp -o ../bin/mgsl_shell.exe",
		"go build go/parse_args.go -o ../bin/arg_parser.exe",
	}

	code := CheckCommandsExist(buildCommands)

	if code != 0 {
		exitWithCode(code)
	}

	// RunCommands(buildCommands)
}
