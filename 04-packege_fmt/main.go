package main

import "fmt"

func main() {
	name := "moawezz"
	fmt.Println("Hello", name)

	a, b := 3, 5
	fmt.Println("sum:", a+b, "divided by 2:", (a+b)/2)

	fmt.Printf("Your age is %d\n", 23)

	fmt.Printf("x is %d , y is %f\n", 12, 3.14)

	fmt.Printf("He says: \"fuck\"\n")

	figure := "cricle"
	radius := 5
	pi := 3.14169

	_ = figure
	_ = pi

	fmt.Printf("Radius: %d\n", radius)
	fmt.Printf("Area: %f\n", float64(radius)*pi*2)

	// %q for qouted strings
	fmt.Printf("This is a %q\n", figure)

	// %v for printing any types
	fmt.Printf("figure: %v, raduis: %v, pi: %v\n", figure, radius, pi)

	// %T for printing out type of variables
	fmt.Printf("figure: %T, raduis: %T, pi: %T\n", figure, radius, pi)

	// %t to format bool values
	opened := true
	fmt.Printf("Is opened: %t\n", opened)

	// %b convert int to binary or base2
	n := 2
	fmt.Printf("Binary of n: %b\n", n)
	// show 8 bits
	fmt.Printf("Binary of n: %08b\n", n)

	// %x decimal to hex
	m := 32
	fmt.Printf("Hex of m: %x\n", m)

	i := 3.2
	j := 6.3

	// display i * j
	fmt.Printf("i * j = %f\n", i*j) // result: 20.160000

	// display only some desimals
	fmt.Printf("i * j = %.3f", i*j) // result: 20.160
}
