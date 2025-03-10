package main

import "fmt"


// jwtToken := "sdfdsf" // this will not work outside of function

const LoginToken string = "sdfdsf" // this will not work outside of function and public becuse starrt captial letter
func main() {
	fmt.Println("variables")

	var username string = "ashish"
	println("username: ", username)
	fmt.Printf("type: %T\n", username)


	var isLoggedIn bool = true
	println("username: ", isLoggedIn)
	fmt.Printf("type: %T\n", isLoggedIn)


	//no var style

	numberofUsers := 34445.4
	println("numberofUsers: ", numberofUsers)
	fmt.Printf("type: %T\n", numberofUsers)

	//multiple variables
	var (
		firstname = "ashish"
		lastname = "raj"
	)
	println("firstname: ", firstname)
	println("lastname: ", lastname)

	// logintoken
	println(LoginToken)
}
