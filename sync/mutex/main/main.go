package main

import (
	"fmt"
	"sync"
)

var x int

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &lock)
	}
	wg.Wait()
	fmt.Printf("final value of x: [%d]", x)
}

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	defer func() {
		m.Unlock()
		wg.Done()
	}()

	m.Lock()
	x = x + 1
}
