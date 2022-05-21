package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // or bufio.NewScanner(os.Open("<file>"))
	for scanner.Scan() {
		if scanner.Text() == "quit" {
			fmt.Println("Quitting")
			os.Exit(3)
		} else {
			fmt.Println("You typed " + scanner.Text())
		}
	}
}
