package main

import (
	"fmt"
	"time"
)

// Create constants for time format
const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutEU  = "2 January, 2006"
)

func main() {
	t := time.Now()

	fmt.Println(t)
	fmt.Println(t.Day(), "--", t.Month(), "--", t.Year())
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format(time.UnixDate))
	fmt.Println(t.Format(time.Kitchen))
	// or inject custom format
	fmt.Println(t.Format("Mon Jan 2 15:04:05 +03 2022"))
	// enter custom date
	startDate := time.Date(2018, 07, 03, 9, 00, 00, 00, time.UTC)

	fmt.Println(startDate)
	fmt.Println(startDate.Format(layoutEU))

	// calculating time spans
	startingDate := time.Date(2021, 07, 03, 9, 00, 00, 00, time.Local)
	elapsed := time.Since(startingDate)
	fmt.Println(elapsed)
	fmt.Printf("Hours: %.2f Minutes: %.2f Seconds: %.0f \n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())

	// Time adding
	today := time.Now()
	future := today.AddDate(0, 6, 0)
	fmt.Println(future.Format(layoutEU))
}
