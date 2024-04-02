package concurrency

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func Foo(wg *sync.WaitGroup) {

	for i := 0; i < 10; i++ {
		fmt.Println("hello there", i)
	}
	wg.Done() // it will decrement the counter
}

func Foo2(wg *sync.WaitGroup) {
	fmt.Println("I am go routine 2")
	time.Sleep(2 * time.Second)

	wg.Done()
}

func Print(num int, gr int, wg *sync.WaitGroup) {
	log.Printf("the go routine %d is printing the value %d", gr, num)
	wg.Done()
}

func SlicePrint(slc []int, wg *sync.WaitGroup) {
	for i := 0; i < len(slc); i++ {
		wg.Add(1)
		go Print(slc[i], i, wg)
	}
}
