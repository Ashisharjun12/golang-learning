package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "hello files"

	file , err := os.Create("./myfile.txt")

	if err != nil {
        panic(err)
    }

length,err := io.WriteString(file , content)

if err != nil {
	panic(err)
}

fmt.Println("content is :" , length)
defer file.Close()
readFile("./myfile.txt")

}


func readFile(filename string) {
	dataByte ,err := ioutil.ReadFile(filename)

	checkNilError(err)

	fmt.Println(string(dataByte))


}

func checkNilError(err error){
	if err != nil {
        panic(err)
    }
}