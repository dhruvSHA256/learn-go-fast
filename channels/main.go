package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// channels are fifo queues
func someFunc(num int, chnl chan<- int) {
	// fmt.Println(num)
	select {
	case chnl <- num:
	}
	// time.Sleep(1)
}

// read from done channel and if value , then return
func doWork(done <-chan bool) {
	// here done is readonly
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("doing work")
		}
	}
	// time.Sleep(1)
}

func fn(x string, chn chan bool) {
	fmt.Println("printing ", x)
	time.Sleep(1 * time.Second)
	// send data to channel
	chn <- true
}
func prod(a []string, chn chan string) {
	for _, x := range a {
		fmt.Println("sending ", x, " to channel")
		chn <- x
	}
	close(chn)
}

func consumer(x string) {
	fmt.Println("Consumeing ", x)
}

func ff(wt *sync.WaitGroup) {
	fmt.Println("hello")
	wt.Done()
}

// var (
// 	count  int
// 	mu     sync.Mutex
// 	wrLock sync.RWMutex
// )

// func increment() {
// 	mu.Lock()
// 	count++
// 	mu.Unlock()
// }

// func read() {
// 	wrLock.RLock()
// 	fmt.Println(count)
// 	wrLock.RUnlock()
// }

var completed bool

func markCompleted() {
	fmt.Println("Marking Completed task")
	completed = true
}
func doTask() bool {
	return 0 == rand.Intn(10)
}
func checkCompleteTask() {
	if completed {
		fmt.Println("Completed task")
	} else {
		fmt.Println("Not completed")
	}
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	iterations := 1000
	// once allow task to completed only once when multiple goroutines are doing it
	var once sync.Once
	var wg sync.WaitGroup
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			if doTask() {
				once.Do(markCompleted)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	checkCompleteTask()

	// locks
	// for i := 0; i < iterations; i++ {
	// 	go increment()
	// 	go read()
	// }
	// time.Sleep(5 * time.Second)
	// fmt.Println("Result: ", count)

	// // wait for goroutines to complete
	// var wt sync.WaitGroup
	// numGoroutines := 3
	// wt.Add(numGoroutines)
	// // wt.Done just subtracts 1 from numGoroutines
	// for i := 0; i < numGoroutines; i++ {
	// 	go ff(&wt)
	// }
	// // if we call wait before done then it will be a deadlock
	// // wt.Wait() just waits till value of numGoroutines is 0
	// wt.Wait()

	// myChannel := make(chan int)
	// myChannel2 := make(chan int)
	// for i := 0; i < 10; i++ {
	// 	go someFunc(i, myChannel)
	// }
	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	myChannel2 <- 69
	// }()
	//     // block till all case run
	// select {
	// case msg := <-myChannel:
	// 	fmt.Println(msg)
	// case msg2 := <-myChannel2:
	// 	fmt.Println(msg2)
	// }
	// for select looop
	// buffered channels
	// myChannel1 := make(chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	select {
	// 	case myChannel1 <- i:
	// 	}
	// }
	// close(myChannel1)
	// for i := range myChannel1 {
	// 	fmt.Println(i)
	// }
	// done channel
	// done := make(chan bool)
	// go doWork(done)

	// time.Sleep(time.Second * 2)
	// close(done)
	// time.Sleep(2)

}
