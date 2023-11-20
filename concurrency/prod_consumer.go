package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Buffer struct {
	// data
	data chan int
	quit chan chan error
}

func (p *Buffer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func producer(p *Buffer) {
	for {
		i := rand.Intn(100)
		delay := rand.Intn(5)
		fmt.Println("Sending to queue: ", i, "with delay of ", delay)
		time.Sleep(time.Second * time.Duration(delay))
		select {
		// add data to queue
		case p.data <- i:
		// if read data from quit channel then close producer
		case quitChan := <-p.quit:
			close(p.data)
			close(quitChan)
			return
		}
	}
}

func consumer(p *Buffer) {
	// read data from queue
	for i := range p.data {
		fmt.Println("Got data from queue:", i)
	}
}

func runProducerConsumer() {
	// create a buffer
	prod := &Buffer{
		data: make(chan int),
		quit: make(chan chan error),
	}
	// run producer
	go producer(prod)
	// run consumer
	consumer(prod)
}
