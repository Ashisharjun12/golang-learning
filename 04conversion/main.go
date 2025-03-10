package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main(){
	fmt.Println("converison")
	fmt.Println("please rate our pizza from 1 to 5: ")

	reader := bufio.NewReader(os.Stdin)

	input , _ := reader.ReadString('\n')

	fmt.Println("thanks  for your rating of ", input)
	fmt.Printf("type of rating is %T\n", input)

	numrating , err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating: ", numrating + 1)
	}



}