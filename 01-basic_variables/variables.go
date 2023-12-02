package main

import "fmt"

func main() {
	// int age = 10; c++ way
	var age int = 10 // Go way
	fmt.Println("age:", age)

	var name = "moawezz"
	_ = name // use this to prevent compile time error.
	// fmt.Println("name:", name)
}
