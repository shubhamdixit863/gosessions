package controllers

func Array() {

	//i := [5]int{}
	//fmt.Println(i)
	//i[0] = 9
	//i[1] = 10
	//i[2] = 11
	//i[3] = 89
	//i[4] = 90
	//
	//fmt.Println(i)
	//df := i[1:3] // slicing syntax
	//fmt.Println("slice from underlying array", df)
	//
	//// slices  are dynamic arrays ,you can assign
	//var slc []int // slc:=[]int{}
	////slc[0] = 9
	//log.Println(slc == nil)
	//slc = append(slc, 789)
	//slc = append(slc, 908)
	//log.Println(slc[1])
	//log.Println(slc == nil)
	//log.Println(slc)

	// creating a slice using make keyword

	//slcmake := make([]int, 9) // is the capacity of the underlying array
	//fmt.Println(slcmake, "length of slice", len(slcmake), "capacity of the slice", cap(slcmake))
	//slcmake[0] = 90
	//fmt.Println(slcmake)
	//
	//slcmake = append(slcmake, 8)
	//fmt.Println(&slcmake[0])
	//fmt.Println(slcmake, "length of slice", len(slcmake), "capacity of the slice", cap(slcmake))
	//slcmake = append(slcmake, 89)
	//fmt.Println(slcmake, "length of slice", len(slcmake), "capacity of the slice", cap(slcmake))

	//fmt.Println(slices.Contains(slcmake, 89))

	// slices comparison  ==

	// the default value of any reference type is nil
	// and slices are reference types
	// and arrays are value types
	// == comparison can only be made on value types not reference types ,reference types can be compared with nil only
	slc := make([]int, 1)
	slc[0] = 9

	slc2 := make([]int, 1)
	slc2[0] = 9

	//fmt.Println("slice comparison", slc ==slc) // not possible
	//fmt.Println("slice comparison with internal package", slices.Equal(slc, slc2))
	//fmt.Println("slice comparison with internal package", reflect.DeepEqual(slc, slc2))

	// arrays are value types
	//arr := [2]int{1, 2}
	//arr1 := [2]int{1, 2}

	//fmt.Println("comparing arrays", arr == arr1)

}

// pass by value --->go is always pass by value  go doesn't have any concept of pass by reference

func SliceReference(slc []int) {
	slc[0] = 90

}

func ArrayReference(slc [1]int) {
	slc[0] = 90

}
