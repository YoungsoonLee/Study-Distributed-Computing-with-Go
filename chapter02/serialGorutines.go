package main

import (
	"fmt"
	"sync"
)

func makeHotelReservation(wg *sync.WaitGroup) {
	fmt.Println("Done making hotel reservation.")
	wg.Done()
}

func bookFlightTickets(wg *sync.WaitGroup) {
	fmt.Println("Done booking flight tickets")
	wg.Done()
}

func orderADress(wg *sync.WaitGroup) {
	fmt.Println("Done ordering a dress.")
	wg.Done()
}

func payCreditCardBills(wg *sync.WaitGroup) {
	fmt.Println("Done paying Credit Card bills.")
	wg.Done()
}

func writeAMail(wg *sync.WaitGroup) {
	fmt.Println("Wrote 1/3rd of the mail.")
	//continueWritingMail1()
	go continueWritingMail1(wg)
}

func continueWritingMail1(wg *sync.WaitGroup) {
	fmt.Println("Wrote 2/3rd of the mail.")
	//continueWritingMail2()
	go continueWritingMail2(wg)
}

func continueWritingMail2(wg *sync.WaitGroup) {
	fmt.Println("Done writing the mail.")
	wg.Done()
}

func listenToAudioBook(wg *sync.WaitGroup) {
	fmt.Println("Listened to 10 minutes of audio book.")
	//continueListeningToAudioBook()
	go continueListeningToAudioBook(wg)
}

func continueListeningToAudioBook(wg *sync.WaitGroup) {
	fmt.Println("Done listening to audio book.")
	wg.Done()
}

var listOfTasks = []func(*sync.WaitGroup){
	makeHotelReservation,
	bookFlightTickets,
	orderADress,
	payCreditCardBills,
	writeAMail,
	listenToAudioBook,
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(listOfTasks))

	for _, task := range listOfTasks {
		go task(&waitGroup) // maxium
	}

	waitGroup.Wait()
}
