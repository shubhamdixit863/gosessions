package controllers

import "cmp"

func Add(a int, b int) int {
	return a + b
}

func AddFloats(a float64, b float64) float64 {
	return a + b
}

// This method becomces the genraric method
func AddAll[T int | float64 | uint, U float64](a T, b T) U {

	return U(a + b)
}

type Num interface {
	int | float64 | uint | string
}

func AddAll2[T Num](a T, b T) T {

	return a + b
}

// Generics with structs

type CustomData interface {
	string | []byte | int
}

type User[T CustomData] struct {
	Name    string
	Age     int
	Address T
}

func Add3[T cmp.Ordered](a, b T) T {
	return a + b

}
