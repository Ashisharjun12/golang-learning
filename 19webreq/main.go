package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


const url = "https://jsonplaceholder.typicode.com/todos/1"
func main() {
	fmt.Println("web requests...")

	// TODO: Send HTTP GET request to the given URL and print the response body
	resposne , err := http.Get(url)
	checkNilError(err)

	defer resposne.Body.Close()

	
content , err :=	ioutil.ReadAll(resposne.Body)

checkNilError(err)
fmt.Printf("type of response %T\n",content)
fmt.Println("data is",string(content))
	
}


func checkNilError(err error){
	if err != nil {
        fmt.Println(err)
        return
    }
}