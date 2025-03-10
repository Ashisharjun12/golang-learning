package main

import "fmt"

func main() {
	fmt.Println("array in golang")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[3] = "Peach"
	fruitlist[2] = "Grapes"

	fmt.Println("fruitlist: ", fruitlist)
	fmt.Println("fruitlist: ", len(fruitlist))


	var vegilist =  [3]string{"Peach", "Grapes", "Apple"}
	fmt.Println("vegilist: ", vegilist)
}