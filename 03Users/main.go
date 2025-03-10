package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")
	// comma ok idiom ||err ok

	input , _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)// prints the input with the newline character



}