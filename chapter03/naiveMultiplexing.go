package main

import "fmt"

func main() {
	channels := [5](chan int){
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
	}

	go func() {
		//starting to wait on channels
		for _, chx := range channels {
			fmt.Println("receiving from ", <-chx)
		}
	}()

	for i := 1; i < 6; i++ {
		fmt.Println("sending on channel: ", i)
		channels[i] <- 1
	}
}
