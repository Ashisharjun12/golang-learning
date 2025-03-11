package main

import "fmt"

func add(a ,b int) int {
	return a + b

}

func getlang() (string, string, string){
	return "a", "b" , "c"
}

func main() {
	res := add(1, 2)

	fmt.Println("adding numbers:", res)

	lang1 , lang2 , _ := getlang()

	fmt.Println("dta lang", lang1 , lang2 )


	// fmt.Println("sum", sums(1,2,3,4,5))

	numbers := []int{1,2,3,4,5}

	result := sums(numbers...)

	fmt.Println("sums num", result)
}


//varadic functions

func sums(nums ...int) int{
	total := 0

	for _ , num := range nums{
		total =total + num
	}

	return total

}