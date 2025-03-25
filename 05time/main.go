package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang")

	presentTime := time.Now()
	// this prints out a date, timestamp, zone, and "margins" (whatever those are)
	fmt.Println(presentTime)

	// this prints just the date. The 01-02-2006 is a documentation specific way to format it
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05"))        // adds the time stamp
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // adds the day of week
	// I'm not sure who wrote this library, but these string formats are straight out of the docs
}
