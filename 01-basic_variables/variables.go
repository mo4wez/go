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
	s := "Go programming language."
	fmt.Println(s)

	// name := "ali" // Error because declared once before.
	name = "yasin" // this is okay.
	name1 := "naseer"
	_ = name1

	// variables multiple declaration
	car, cost := "Pars", 750
	fmt.Println(car, cost)

	// car, cost := "405", 530 // Cause Error
	car, year := "405", 1393 // atleast one new variable on the left side of :=
	_ = year

	var opened = false
	opened, file := true, "names.txt"
	_, _ = opened, file

	// another way to declare multipe variables for good readability
	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	// another way
	var x, y, z int
	fmt.Println(x, y, z)

	// multiple assignment
	var i, j int
	i, j = 1, 3
	fmt.Println("before swap i, j:", i, j)

	// swap variables
	j, i = i, j
	fmt.Println("after swap i,j:", i, j)

	// use expression in :=
	sum := 5 + 3.7
	fmt.Println("sum:", sum)
}
