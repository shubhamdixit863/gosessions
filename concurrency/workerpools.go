package concurrency

import "sync"

type Job struct {
	Number int // our job to multiply the number in Job struct by 2
}

type Result struct {
	Data int
}

func Worker(jobChannel chan Job, resultChannel chan Result, wg *sync.WaitGroup) {
	// this worker will do some job
	data := <-jobChannel
	result := data.Number * 2
	//log.Println(result)
	resultChannel <- Result{result} /// this is will also be a blocking job
	wg.Done()

}
