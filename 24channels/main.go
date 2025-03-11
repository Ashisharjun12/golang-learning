package main

import (
	"fmt"
	"time"
	// "math/rand/v2"
)

// func processNum(numChan chan int){

// 	for num := range numChan{
// 		fmt.Println("processing number ",  num) //recive  number

// 	}

// }


func Sum(result chan int , num1 int , num2 int){
	result <- num1 + num2

}

//go routine syncronizer
func task(done chan bool ){
	defer func(){done <- true}()
	fmt.Println("processing task...")
}

func emailSender(emailChan chan string){
	defer close(emailChan)  //close the channel after sending all data becuse to prevent deadlock
	for email := range emailChan{
        fmt.Println("sending email to ", email)
    }
}

func main() {

	//multiple channels and commmunication betwwwn channels
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
        chan1 <- 10
        chan1 <- 20
    }()

	go func(){
		chan2 <- "string data on chan2"

	}()

	for i := 0 ; i<2 ;i++ {
		select{
		case chan1Value := <-chan1:
			fmt.Println("received from chan1: ", chan1Value)

		case chan2Value := <-chan2:
			fmt.Println("received from chan2: ", chan2Value)

		}
	
	}




	//these are bufferd channels
	emailChan := make(chan string ,100) // buffer channel 100 data without blaock 
	
	go emailSender(emailChan) //implemmet queqe sysytem

	for i:=0 ; i<10 ; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
		time.Sleep(time.Second)
	}

	fmt.Println("done...")


	// emailChan <- "1@gmail.com"
	// emailChan <- "2@gmail.com"


	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)






	//these are unbufferd channels

	done := make(chan bool)

	go task(done)

	<- done //  dat recived  and blocking  and wait until task complete



	result := make (chan int)

	go Sum(result, 4 , 5)

	res := <- result //  dat recived  and blocking 

	fmt.Println("Sum : ", res)



	// numChan := make(chan int)

	// go processNum(numChan)
 
	// // numChan <- 5 // send number

	// for {
	// 	numChan <- rand.IntN(100)
	// }

	

	// messageChan := make(chan string)

	// messageChan <- "ping , send data to message channel" //send data to message channel  and blocking

	// msg := <-messageChan // recive data from message channel

	// fmt.Println("receive data from message channel", msg)
}