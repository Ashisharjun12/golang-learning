package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task", id)
}

func main() {
	for i :=0 ;i<=10 ; i++{
		// task(i)

		go task(i) // run in go routine new lightweigh threads and non-blocking threads so order is random  cpu muti thread use
	}

	time.Sleep(time.Second *2)
}