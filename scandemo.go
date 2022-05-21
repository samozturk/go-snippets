package main

import "fmt"

func main() {
	var name string
	fmt.Println("enter your name")
	inputs, _ := fmt.Scanf("%s", &name)
	fmt.Printf("Hello %v, inputs: %d", name, inputs)
}
