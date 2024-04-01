package controllers

import (
	"log"
	"reflect"
	"strconv"
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
	// reflect package
}

// Switch case based type assertions
// Better way to handle the type assertions in golang

func AnyKind2(param interface{}) {

	switch d := param.(type) {
	case string:
		log.Println(strings.ToUpper(d))

	case int:
		log.Println(d + 9)

	case bool:
		log.Println(d)

	default:
		log.Println("Unknown type")

	}

}

// type conversion  is when you convert one type into another

func TypeConversion() {

	// if you want to convert numeric types into numeric types
	// from int to float or vice versa

	// you can use the constructor method
	i := 9 // int
	// convert it to floa
	log.Println(reflect.TypeOf(float64(i)))
	j := 8.9
	log.Println(reflect.TypeOf(int(j)))
	// number into string
	atoi, err := strconv.Atoi("89")
	log.Println(atoi)
	if err != nil {
		log.Println(err)
		return
	}

	// strigg in go are basically the byte slice
	str := "he"
	log.Println([]byte(str))
}
