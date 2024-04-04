package concurrency

import (
	"log"
	"sync"
)

type Job struct {
	Number int // our job to multiply the number in Job struct by 2
}

type Result struct {
	Data int
}

func Worker(jobChannel chan Job, resultChannel chan Result, jobId int, wg *sync.WaitGroup) {
	// this worker will do some job

	data := <-jobChannel
	result := data.Number * 2
	log.Println("Job id", jobId)
	resultChannel <- Result{result} /// this is will also be a blocking job
	wg.Done()
}
