package main

import (
	"fmt"
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
func main() {
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
