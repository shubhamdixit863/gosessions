package controllers

import "fmt"

func Maps() {
	// ways to create a map
	// maps are also reference types
	var m map[string]string
	fmt.Println(m)
	fmt.Println(m == nil)
	//	m["name"] = "shubham"  // this won't work because map is nil

	// how to solve this issue
	mp := make(map[string]string)
	mp["name"] = "now it will work"
	fmt.Println(mp)

	// other syntax
	mk := map[string]string{
		"key": "value",
	}
	fmt.Println(mk)
	fmt.Println(mk == nil)
	mk["name"] = "oo"
	fmt.Println(mk)

	// map keys can be anything which can be compared int ,string ,array
	// but it wont be a slice

	//mj := make(map[[2]int]string)
	//mj[[2]int{90, 9}] = "hey there"
	//fmt.Println(mj)
}

func MapModify(mp map[string]string) {
	mp["key"] = "data"
}
