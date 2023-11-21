package main

import (
	"fmt"
	"sync"
	"time"
)

var philosopher = []int{1, 2, 3, 4, 5}
var wg sync.WaitGroup
var delay = 1 * time.Second

func eat(p int, l, r *sync.Mutex) {
	defer wg.Done()
	fmt.Println(p, " will not be eating")
	for i := 0; i < 3; i++ {
		// lock
		fmt.Println("\t", p, " Picking right fork")
		r.Lock()
		fmt.Println("\t", p, " Picking left fork")
		l.Lock()
		// eat
		fmt.Println("\t", p, " Eating")
		time.Sleep(delay)
		// think
		fmt.Println(p, " is thinking")
		time.Sleep(delay)
		// unlock
		fmt.Println("\t", p, " Releasing right fork")
		r.Unlock()
		fmt.Println("\t", p, " Releasing left fork")
		l.Unlock()
		time.Sleep(delay)
	}
	fmt.Println(p, " is done eating")
}

func runDiningPhilosopher() {
	wg.Add(len(philosopher))
	l := &sync.Mutex{}
	for _, p := range philosopher {
		r := &sync.Mutex{}
		go eat(p, l, r)
		l = r
	}
	wg.Wait()
}
