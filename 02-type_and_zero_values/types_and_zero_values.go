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
}
