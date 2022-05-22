package main

import (
	"fmt"
	"log"
)

type messageType int

// Declaring multiple constants
const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
)
const (
	InfoColor    = "\033[1;32m%s\033[0m" // Green
	WarningColor = "\033[1;33m%s\033[0m" // Yellow
	ErrorColor   = "\033[1;31m%s\033[0m" // Red
	// Refer to https://www.shellhacks.com/bash-colors/ for different colors.
)

func showMessage(messageType messageType, message string) {
	switch messageType {
	case INFO:
		printMessage := fmt.Sprintf("\nInformation: \n%s\n", message)
		log.Printf(InfoColor, printMessage)
	case WARNING:
		printMessage := fmt.Sprintf("\nWarning: \n%s\n", message)
		fmt.Printf(WarningColor, printMessage)
	case ERROR:
		printMessage := fmt.Sprintf("\nError: \n%s\n", message)
		fmt.Printf(ErrorColor, printMessage)
	}
}

func main() {

	showMessage(ERROR, "oops")
	showMessage(INFO, "know")
	showMessage(WARNING, "hello")
}
