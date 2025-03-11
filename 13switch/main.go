package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Printing the string
	fmt.Println("learning switch case...")

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10) + 1
	fmt.Printf("Random number: %d\n", randomNumber)

	switch randomNumber {
		case 1, 3, 5, 7, 9:
            fmt.Println("The number is odd.")
			fallthrough
        case 2, 4, 6, 8:
            fmt.Println("The number is even.")
        default:
            fmt.Println("The number is neither odd nor even.")
	}
}