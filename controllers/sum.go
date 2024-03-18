package controllers

import "fmt"

var C = 9 // automatically eexported from the package
var f = 9 // you cant use it outside the package

// this is an exported function and It can be used outside this package as well

func Sum(name string, address string, age int) { // parameters  -- parameterName paramtertype
	fmt.Println(name, address, age)
}

// Function with return value

func Addition(a, b int) int {

	sum := a + b
	return sum
}

// You can return multiple values from the function as well

func AdditionSubtraction(a, b int) (int, int) {

	sum := a + b
	diff := a - b
	return sum, diff
}

// you can return n number of values from the function

//func AdditionSubtraction2(a, b int) (int, int, int, int, int) {
//
//	sum := a + b
//	diff := a - b
//	return sum, diff
//}

// defer

// variadic function

func Foo(a ...int) {

	fmt.Println(a) // it prints slice data structure

}
