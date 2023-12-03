package main

import "fmt"

var taskDone bool

func main() {
	fmt.Println("Naming Conventions")

	// Best practice
	var mv int // max value
	_ = mv
	// var max_value int // Not Ok

	// in small scope
	// var packegesReceived int // Not Ok, too long
	// var b int // Good practice

	// const MAX_VALUE = 100 // Not Ok
	// const N = 1000 // better

	// maxValue := 1000 // recommended
	// max_value := 1000 // not idiometic

	// writeToDB := true // ok, idiometic
	// writeToDb := true // not idiometic

}
