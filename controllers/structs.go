package controllers

//
//import "fmt"
//
//// structs are composite types
//
//type Student struct {
//	Name    string
//	Email   string
//	Address string
//	Age     int
//	Exam    // composition or embediing
//}
//
//type Student2 struct {
//	Name    string
//	Email   string
//	Address string
//	Age     int
//	Exam    Exam // composition or embediing
//}
//
//// go doesnt have inheritance --> is a relationship Animal class and Dog class
//// instead go have composition  --> has a relationship
//
//type Exam struct {
//	Name string
//	Date string
//}
//
//func StructFun() {
//	c := Student{
//		Name:    "Shubham",
//		Email:   "sd@ymail.com",
//		Address: "delhi",
//		Age:     20,
//		Exam:    Exam{Name: "Physics", Date: "9-09-2010"},
//	}
//
//	fmt.Println(c.Name)
//	fmt.Println(c.Exam.Name)
//
//	c2 := Student2{
//		Name:    "Shubham",
//		Email:   "sd@ymail.com",
//		Address: "delhi",
//		Age:     20,
//		Exam:    Exam{Name: "Physics", Date: "9-09-2010"},
//	}
//
//	//fmt.Println(c2.Type)  // not valid
//	fmt.Println(c2.Exam.Name)
//
//}
//
///// Food delivery app
//
////class User{
////	Nid string
////	PermannetAddress string
////	Name string
////}
////
////
////class Customer extends User{
////
//// function foo(){
////}
////	DeliveryAddress :[]string
////	FoodOrdered:[]string
////
////}
////
////class DeliveryBoy extends User{
////	HasAVehicle bool
////	MonthlYPayment bool
////}
////
////type Customer struct {
////	CommonInfo
////}
////
////type DeliveryBoy struct {
////	CommonInfo
////}
////
////type CommonInfo struct {
////	Nid string
////	PermannetAddress string
////	Name string
////}
//
//// If any function belongs to an object we call it as method
///*
//type Human struct {
//	Name string
//	Age  int
//}
//
// */
//
//// value receiver method
//func (h Human) Walk() {
//	fmt.Println(" method invoked")
//	h.Name = "not a new john"
//}
//
//// pointer receiver method
//
//// value receiver method
//func (h *Human) Sing() {
//	fmt.Println("singing  method invoked")
//	h.Name = "New john"
//}
