package main

import (
	"fmt"
	"time"
)

func main() {

	worker := NewWorker()

	go func() {
		fmt.Println("wait for terminate")
		<-worker.StopC
		fmt.Println("terminated")

	}()

	go func() {

		for i := 0; i < 10; i++ {
			fmt.Println(i, worker.Msg)
			time.Sleep(1000)
		}

		close(worker.StopC)
	}()

	for {

	}
}

type Worker struct {
	StopC  chan int
	StopCS []chan int
	Msg    string
}

func NewWorker() *Worker {
	worker := new(Worker)
	worker.StopC = make(chan int)
	worker.StopCS = make([]chan int, 0, 3)
	worker.Msg = "Hello"
	return worker
}
