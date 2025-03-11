package main

import (
	"fmt"
	"sync"
	
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		// task(i)
		wg.Add(1)
		go task(i, &wg) // run in go routine new lightweigh
		//threads and non-blocking threads so order is random  cpu muti thread use
	}

	wg.Wait()

}
