package main

import (
	"fmt"
	"os"

	"gosession/controllers"
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
	////const name = "shubham"  // you cannot change the values assigned to const variable
	////name="sjsjjs"
	//// Variable creation with var keyword
	//// first way
	////var data string
	//
	//// second way
	//
	////var data2 = "hello people"
	//
	//// third way you donot need to use the var keyword
	//// short hand syntax
	////data3 := "something else" // this automatically resolves the type
	////data3 = 9  // static type checking
	//
	////fmt.Println(reflect.TypeOf(data3))
	////fmt.Println(name)
	//
	////var c int
	////var d bool
	////var e float32
	////var f string
	//
	//fmt.Println(c, d, e, f) // int string bool float
	//
	//// conversion of one data type to another
	//
	//numInt := 9
	//
	//log.Println(reflect.TypeOf(numInt))
	//numFloat := float32(numInt)
	//log.Println(reflect.TypeOf(numFloat))
	//
	//stringMy := "23"
	//log.Println(strconv.Atoi(stringMy))//fmt.Println(reflect.TypeOf(data3))
	//fmt.Println(name)
	//
	//var c int
	//var d bool
	//var e float32
	//var f string
	//
	//fmt.Println(c, d, e, f) // int string bool float
	//
	//// conversion of one data type to another
	//
	//numInt := 9
	//
	//log.Println(reflect.TypeOf(numInt))
	//numFloat := float32(numInt)
	//log.Println(reflect.TypeOf(numFloat))
	//
	//stringMy := "23"
	////log.Println(strconv.Atoi(stringMy))	log.Println(reflect.TypeOf(numInt))
	////numFloat := float32(numInt)
	////log.Println(reflect.TypeOf(numFloat))
	////
	////stringMy := "23"
	////log.Println(strconv.Atoi(stringMy))//fmt.Println(reflect.TypeOf(data3))
	////fmt.Println(name)
	////
	////var c int
	////var d bool
	////var e float32
	////var f string
	////
	////fmt.Println(c, d, e, f) // int string bool float
	////
	////// conversion of one data type to another
	////
	////numInt := 9
	////
	////log.Println(reflect.TypeOf(numInt))
	////numFloat := float32(numInt)
	////log.Println(reflect.TypeOf(numFloat))
	////
	////stringMy := "23"
	////log.Println(strconv.Atoi(stringMy))

	// defer is used to execute the instruction just before the function returns
	//defer fmt.Println("hii there  ---------")
	//
	////fmt.Println(controllers.C, controllers.C)
	//controllers.Sum("john", "Delhi", 23) // pass the arguments
	//
	//value := controllers.Addition(2, 3)
	//fmt.Println(value)
	//
	//value1, value2 := controllers.AdditionSubtraction(8, 7)
	//
	//fmt.Println("Two returned values", value1, value2)

	//foo()
	// you can pass n number of arguments
	controllers.Foo(2, 3, 4, 5, 6, 6, 7, 8, 9, 10)

}

func infinite() {
	go func() {
		for {

		}
	}()

}

func foo() {
	open, err := os.Open("go.mod")
	if err != nil {
		return
	}
	defer open.Close()

	defer fmt.Println("1")

	defer fmt.Println("3")

	fmt.Println("hi there")
	fmt.Println("hey people")

}
