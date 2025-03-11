package main

import "fmt"

func main(){
	fmt.Println("studying structs")
	//no inheritance , no super , no parents


	user := User{Name: "Ashish", Email: "ashish@gmail.com", Status: true, Age: 25}
	fmt.Println(user)
	//single value
	fmt.Println("name is ", user.Name , " and email is ", user.Email)

}


type User struct {
	Name string
	Email string
	Status bool
	Age int
}