package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}() // This is an anonymous function

	go func() {
		defer waitGrp.Done()
		fmt.Println("World")
	}()
	waitGrp.Wait()
}
