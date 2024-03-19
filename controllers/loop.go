package controllers

import (
	"fmt"
	"log"
)

// for keyword
// which you can use anykind of loop may while loop or typical while loop too

func GenerateNum() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

}

// while loop
func GenerateNumWhile() {
	i := 0
	for i < 10 {
		// some operation here
		fmt.Println("hello there")

		i++
	}
}

// infinite while loop

func GenerateNumWhileInfinite() {
	i := 0
infiniteloop: // label to the loop
	for {
		// some operation here
		fmt.Println("hello there")
		if i == 100 {
			break infiniteloop
		} else {
			i++
			continue

		}

	}
}

func GenerateNumWhileInfiniteTwo() {
	i := 0
infiniteloop: // label to the loop
	for {
		for {
			// some operation here
			fmt.Println("hello there")
			if i == 100 {
				break
			} else if i == 200 {
				break infiniteloop
			} else {
				i++
				continue

			}

		}

	}
	log.Println("Broke from outer loop ,value of i ", i)
}

func SwitchCase(fruit string) string {

	switch fruit {
	case "apple":
		return "red"
	case "banana":
		return "yellow"
	case "kiwi":
		return "green"

	default:
		return "No matching expression"

	}

}
