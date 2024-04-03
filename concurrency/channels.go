package concurrency

import (
	"log"
	"time"
)

func SendTheData(ch chan int) {
	log.Println("hi")
	for i := 0; i < 10; i++ {
		log.Println("hiiii from the sender")
	}

	/// It will send the data to the channel
	ch <- 90 // this would be a blocking operation if the channel is unbuffered
	log.Println("furher from send the data routine ")
	log.Println("furher from send the data routine ")
	log.Println("furher from send the data routine ")
	log.Println("furher from send the data routine ")
	for i := 0; i < 9; i++ {
		log.Println("hekk")
	}

}
func Receicvethedata(ch chan int, ch2 chan struct{}) {

	time.Sleep(10 * time.Second)
	log.Println("the data from the sender", <-ch)

	// i will send the data to this channle ch2

	ch2 <- struct{}{}

}
