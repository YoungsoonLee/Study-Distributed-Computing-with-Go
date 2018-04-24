package main

import (
	"fmt"
	"sync"
	"time"
)

func printTime(msg string) {
	fmt.Println(msg, time.Now().Format("15:04:05"))
}

func writeMail1(wg *sync.WaitGroup) {
	printTime("Done writing mail #1.")
	wg.Done()
}

func writeMail2(wg *sync.WaitGroup) {
	printTime("Done writing mail #2.")
	wg.Done()
}

func writeMail3(wg *sync.WaitGroup) {
	printTime("Done writing mail #3.")
	wg.Done()
}

func listenForever() {
	for {
		printTime("Listening...")
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	go listenForever()

	time.Sleep(time.Nanosecond * 10)

	go writeMail1(&waitGroup)
	go writeMail2(&waitGroup)
	go writeMail3(&waitGroup)

	waitGroup.Wait()
}
