package main

import "fmt"

func main() {
	// int age = 10; c++ way
	var age int = 10 // Go way
	fmt.Println("age:", age)

	var name = "moawezz"
	_ = name // use this to prevent compile time error.
	fmt.Println("name:", name)

	// variables short declaration
	s := "Go programminig language."
	fmt.Println(s)

	// name := "ali" // Error because declared once before.
	name = "yasin" // this is okay.
	name1 := "naseer"
	_ = name1

}
