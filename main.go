package main

import (
	"fmt"
	"log"

	"gosession/concurrency"
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

	//controllers.AnyKind2(89)
	//controllers.AnyKind2("shello")
	//controllers.AnyKind2(true)
	//controllers.AnyKind2(89.9)
	//controllers.AnyKind2(controllers.Lfu{})

	//controllers.TypeConversion()
	//fmt.Println(controllers.Add3(8, 9))
	//fmt.Println(controllers.Add3(8.9, 9.7))
	//fmt.Println(controllers.Add3("hello", "world"))

	// Concurrency
	//var wg sync.WaitGroup

	// this will make sure that our main routine waits for atleast 1 go routine
	//wg.Add(2)

	//// you use make keyword to create a channel
	//ch := make(chan int) // unbuffered channel
	//blockMainRoutineChannel := make(chan struct{})
	//
	//go concurrency.SendTheData(ch)
	//go concurrency.Receicvethedata(ch, blockMainRoutineChannel)
	//
	//go func() {
	//	time.Sleep(10 * time.Second)
	//	log.Println("I am an anonymous function")
	//	wg.Done()
	//}()
	//slc := []int{8, 9, 7, 70, 6, 45}
	//concurrency.SlicePrint(slc, &wg)
	//<-blockMainRoutineChannel
	//log.Println("Data received in main routine", <-ch)
	// you have to stop the main routine from exiting
	// wait groups

	//wg.Wait() // it will block the main routine until the counter is 0

	// here we will create job channel and result channels

	// create a job channel
	//jobC := make(chan concurrency.Job, 5) // it as a  buffered channel because we dnt want to block it
	//
	//resultChannel := make(chan concurrency.Result)
	//
	//jobs := []concurrency.Job{
	//	{Number: 1},
	//	{Number: 10},
	//	{Number: 20},
	//	{Number: 11},
	//	{Number: 34},
	//}
	//
	//// we will spawn our workers
	//var wg sync.WaitGroup
	//
	//for i := 0; i < 5; i++ {
	//	go concurrency.Worker(jobC, resultChannel, i, &wg)
	//}
	//
	//// we have to send jobs to channel
	//
	//for i := 0; i < len(jobs); i++ {
	//	wg.Add(1)
	//	jobC <- jobs[i] // it wont be a blocking operation as it is the buffered channel
	//}
	//close(jobC)
	//
	//// its a blocking operation for main go routine
	//
	//go func() {
	//	wg.Wait()
	//	close(resultChannel)
	//}()
	//
	//for res := range resultChannel {
	//	log.Println(res)
	//}
	//slc := []int{8, 2, 11, 5, 3, 20, 1}
	//ch := make(chan int)
	//fmt.Println(concurrency.BubbleSortConcurrent(slc, ch))

	innerCh := make(chan int) // we will be passing this to the outerLoop
	exitch := make(chan struct{})
	slc := []int{8, 2, 11, 5, 3, 20, 1}
	go concurrency.OuterLoop(slc, innerCh, exitch)

	// next thing is we will listen to the  channels or wait for the data to come to channels
loop:
	for {

		select {
		case i := <-innerCh:
			concurrency.InnerLoop(i, slc)
		case <-exitch:
			break loop

		}

	}
	fmt.Println("sorted slice", slc)
}

func OnlyForHumans(human controllers.Human) {
	log.Println("congratulations you are human", human)
}

// selection sort with same technique --------

// we will create a tcp socket server   // net package

// mutexes

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

// try to do bubble sort with go routines
