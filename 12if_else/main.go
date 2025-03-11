package main

import "fmt"
func main() {
	fmt.Println("if else studing...")

	loginCout :=23
	var result string

	if loginCout <10 {
		result = "regular user"

	} else if loginCout > 10{
		result = "admin user"    
	} else {
		result = "invalid user"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	}else {
		fmt.Println("Number is odd")
	}

	//weired syntex

	if num := 3 ; num >4{
		fmt.Println("Number is greater than 4")
	}else{
		fmt.Println("Number is not greater than 4")
	}
  

	//

	// if err != nil{
	// 	fmt.Println("Error: ", err)
	// }
}