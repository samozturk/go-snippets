package main

import (
	"flag"
	"fmt"
)

func main() {
	archPtr := flag.String("arch", "x86", "CPU Type")
	flag.Parse()

	switch *archPtr {
	case "x86":
		fmt.Println("32 bit mode")
	case "AMD64":
		fmt.Println("64 bit mode")
	case "IA64":
		fmt.Println("remember IA64?")

	}
}
