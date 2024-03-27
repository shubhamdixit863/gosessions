package controllers

import "log"

// Whole point of interface is to impose some behaviour
type Human interface {
	Friendship()
	Talk()
}

type Animal struct {
	Legs  int
	Tail  int
	Hands int
}

type Bird struct {
	Feathers int
}

// What we do implement the interface

func (an Animal) Friendship() {
	log.Println("Lets be friends")
}

func (an Animal) Talk() {
	log.Println("Lets be friends")
}

func (b Bird) Friendship() {

}

func (b Bird) Talk() {

}

type PizzaHut interface {
	SellFarmhousePizza()
	PizzaHutLogo()
	StaffWearPizzaHutDress()
}

type SajidAvenue struct {
}
