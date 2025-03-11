package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

// increaseViews increases the post's views by 1, ensuring thread safety with a mutex. and solve race conditions, lock where modification where only
func (p *post) increaseViews(wg *sync.WaitGroup) {

	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1

}

func main() {
	var wg sync.WaitGroup
	mypost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mypost.increaseViews(&wg)

	}

	wg.Wait()
	fmt.Println("inc post", mypost.views)

}
