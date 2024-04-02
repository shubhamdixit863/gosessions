package concurrency

import (
	"log"
	"sync"
	"time"
)

func Sender(ch chan int, wg *sync.WaitGroup) {
	// send some data to the channel
	log.Println("Sender here --- trying to send data------")
	time.Sleep(5 * time.Second)
	ch <- 8
	wg.Done()
}

func Receiver(ch chan int, wg *sync.WaitGroup) {
	// will receive the data from the channel
	log.Println("I have  reached the reciving location waiting for you to send")
	c := <-ch
	log.Println("Now we have recived the value")

	log.Println(c)
	wg.Done()
}
