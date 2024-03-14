package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

// Two keywords ,const ,var
// var
// multiple variable declartion with var keyword
var (
	name    = "john"
	address = "delhi"
)

// multiple variables with const
const (
	one = 1
	two = "two"
)

func main() {
	//const name = "shubham"  // you cannot change the values assigned to const variable
	//name="sjsjjs"
	// Variable creation with var keyword
	// first way
	//var data string

	// second way

	//var data2 = "hello people"

	// third way you donot need to use the var keyword
	// short hand syntax
	//data3 := "something else" // this automatically resolves the type
	//data3 = 9  // static type checking

	//fmt.Println(reflect.TypeOf(data3))
	fmt.Println(name)

	var c int
	var d bool
	var e float32
	var f string

	fmt.Println(c, d, e, f) // int string bool float

	// conversion of one data type to another

	numInt := 9

	log.Println(reflect.TypeOf(numInt))
	numFloat := float32(numInt)
	log.Println(reflect.TypeOf(numFloat))

	stringMy := "23"
	log.Println(strconv.Atoi(stringMy))

}
