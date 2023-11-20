package main

import (
	"fmt"
	"sync"
	"time"
)

// main is a goroutine
// all goroutines are managed by go scheduler

func pp(s any, wg *sync.WaitGroup) {
	// wg.Done decreases counter
	defer wg.Done()
	fmt.Println(s)
}
func main() {
	var wg sync.WaitGroup
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	time.Sleep(1 * time.Second)
	// one entry for each goroutine
	// wg.Add init counter
	wg.Add(len(arr))

	for _, x := range arr {
		// to wait for completing all goroutines we will use waitgroup
		go pp(x, &wg)
	}
	// wg.Wait waits till counter gets to 0
	wg.Wait()
}
