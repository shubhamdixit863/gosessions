package controllers

import (
	"log"
	"strings"
)

func Bar(name string, age int) {

}

// Lets say you want to make your function any kind of type

func AnyKind(param interface{}) {
	// What if you want to perform some operation
	//c:=param+9  // not valid first convert it to the integer type
	// you do type assertion

	c, ok := param.(int)
	if ok {
		log.Println(c + 9)

	}
	// we will assert for string
	d, ok := param.(string)
	if ok {
		log.Println(strings.ToUpper(d))

	}

	// switch case based type assertion that

}
