package main

import (
	"fmt"
	"os/exec"
	"strconv"
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

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString, fmt.Sprint(args...))
	}
	return sprint
}

func CommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func CheckCommandsExist(commands []string) {
	var command_exists bool
	var printColor func(...interface{}) string

	for _, command := range commands {
		command_exists = CommandExists(command)

		if command_exists {
			printColor = Success
		} else {
			printColor = Fatal
		}

		fmt.Printf("%s\t| %s\n", command, printColor(strconv.FormatBool(command_exists)))
	}
}

func main() {
	commands := []string{"go", "g++"}

	fmt.Println("Checking command existence...")

	CheckCommandsExist(commands)
}
