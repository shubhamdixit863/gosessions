package concurrency

import "fmt"

func Foo() {

	for i := 0; i < 10; i++ {
		fmt.Println("hello there", i)
	}

}
