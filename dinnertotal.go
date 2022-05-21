package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usage: dinnertotal <Total Meal Amount> <Tip Percentage>")
		fmt.Println("Example: dinnertotal 20 10")
	} else {
		if len(args) != 2 {
			fmt.Println("You must enter arguments!, type /help for help.")
		} else {
			mealTotal, _ := strconv.ParseFloat(args[0], 32)
			tipAmount, _ := strconv.ParseFloat(args[1], 32)
			fmt.Printf("Meal %.2f, tip %.2f \n", float32(mealTotal), float32(tipAmount))
		}
	}
}
