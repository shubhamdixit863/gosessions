package controllers

import (
	"fmt"
	"log"
)

func Array() {

	i := [5]int{}
	fmt.Println(i)
	i[0] = 9
	i[1] = 10
	i[2] = 11
	i[3] = 89
	i[4] = 90

	fmt.Println(i)
	df := i[1:3] // slicing syntax
	fmt.Println("slice from underlying array", df)

	// slices  are dynamic arrays ,you can assign
	var slc []int // slc:=[]int{}
	//slc[0] = 9
	log.Println(slc == nil)
	slc = append(slc, 789)
	slc = append(slc, 908)
	log.Println(slc[1])
	log.Println(slc == nil)
	log.Println(slc)

	// creating a slice using make keyword

	slcmake := make([]int, 9, 90) // is the capacity of the underlying array
	//log.Println(slcmake == nil)
	slcmake[0] = 9
	//slcmake[12] = 90

}
