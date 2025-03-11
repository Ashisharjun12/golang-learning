package main

import (
	"fmt"
	
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	some :=[]string{"java" , "python" , "numpy" }

	printSlice(nums)
	printSlice(some)

}

func printSlice[T int | string](items []T) {
	for _, items := range items {
		fmt.Println(items)
	}

}