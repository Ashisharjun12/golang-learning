package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time in Go")

	presentTime := time.Now()

	fmt.Println("Current time is: ", presentTime)
	fmt.Println("currnet time is :" ,presentTime.format("01-02-2006 15:04:05"))

}