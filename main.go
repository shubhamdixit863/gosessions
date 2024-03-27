package main

import (
	"log"

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
	//controllers.Foo(2, 3, 4, 5, 6, 6, 7, 8, 9, 10)
	//controllers.GenerateNum()
	//controllers.GenerateNumWhile()
	//controllers.GenerateNumWhileInfinite()
	//fmt.Println(controllers.SwitchCase("banana"))
	//fmt.Printf("For banana the colors is %s \n", controllers.SwitchCase("banana"))
	//controllers.GenerateNumWhileInfiniteTwo()
	//controllers.Array()
	//slc := make([]int, 1)
	//slc[0] = 23
	//fmt.Println("Before", slc)
	//controllers.SliceReference(slc)
	//fmt.Println("After", slc)
	//
	//arr := [1]int{1}
	//
	//fmt.Println("Before Array", arr)
	//controllers.ArrayReference(arr)
	//fmt.Println("After Array", arr)

	//controllers.Maps()
	//mp := make(map[string]string)
	//mp["name"] = "john"
	//fmt.Println(mp)
	//controllers.MapModify(mp)
	//fmt.Println(mp)

	//controllers.Pointers()
	//
	//d := 9
	//pointerVar := &d
	//controllers.PointerArgs(pointerVar)

	//controllers.StructFun()

	//john := controllers.Human{
	//	Name: "John",
	//	Age:  20,
	//}
	//
	//john.Walk()
	//
	//fmt.Println(john)
	//john.Sing()

	// pointer value of object

	//john2 := &controllers.Human{
	//	Name: "John2",
	//	Age:  22,
	//}
	//
	////john2.Walk()
	//
	//john2.Sing()
	//fmt.Println(john2)

	//shubham := controllers.Animal{
	//	Legs:  2,
	//	Tail:  0,
	//	Hands: 2,
	//}
	//// Is shubham a human ?
	//shubham.Friendship()
	//
	//bird := controllers.Bird{Feathers: one}
	//
	//OnlyForHumans(shubham)
	//OnlyForHumans(bird)

	//lfu := &controllers.Lfu{}
	//cache := controllers.InitCache(lfu)
	//cache.Add("a", "1")
	//cache.Add("b", "7")
	//cache.Add("c", "9")
	//
	//cache.SetEvictionAlgo(&controllers.Lru{})
	//
	//cache.Add("a", "1")
	//cache.Add("b", "7")
	//cache.Add("c", "9")

	controllers.AnyKind(89)
	controllers.AnyKind("shello")
	controllers.AnyKind(true)
	//controllers.AnyKind(89.9)
	//controllers.AnyKind(controllers.Lfu{})

}

func OnlyForHumans(human controllers.Human) {
	log.Println("congratulations you are human", human)
}

//func infinite() {
//	go func() {
//		for {
//
//		}
//	}()
//
//}
//
//func foo() {
//	open, err := os.Open("go.mod")
//	if err != nil {
//		return
//	}
//	defer open.Close()
//
//	defer fmt.Println("1")
//
//	defer fmt.Println("3")
//
//	fmt.Println("hi there")
//	fmt.Println("hey people")
//
//}
