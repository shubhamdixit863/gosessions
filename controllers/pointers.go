package controllers

import "fmt"

func Pointers() {
	//var c int
	f := 9

	d := &f // d is a pointer to the variable f
	fmt.Println(d)

	// referencing and dereferencing
	//fmt.Println(d)

	// you can also access the variable using the pointer by dereferencing
	fmt.Println(*d) // it will print out 9

}

// how we can pass pointer to the function

func PointerArgs(num *int) {
	fmt.Println("pointer var", num)    // pointer value the address
	fmt.Println("pointer value", *num) // pointer value the address

}

func PointerForSlice() *[]int {
	k := make([]int, 1)

	// create pointer to the slice

	//l:=[5]int{1,2,3,3,4}
	//j:=&l

	h := &k
	return h
}
