package main

import "fmt"

func main() {
	fmt.Println("types and zero values.")

	// type inferring
	var a = 2
	var b = 4.2

	// Error: can not assign type float to int
	// a = b
	a = int(b) // it's ok
	fmt.Println(a, b)

	// Error: go is a strong typed language
	// var x int
	// x = "5"

	// for numeric type: 0
	// for string type: ""
	// for bool type: false
	// for pointer type: nil

	var x int
	var grade float64
	var name string
	var done bool

	fmt.Println(x, name, grade, done)
}
