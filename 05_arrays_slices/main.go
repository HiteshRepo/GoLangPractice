package main

import "fmt"

func main() {
	// Arrays
	// var fruitArray [2]string

	// Assign values
	// fruitArray[0] = "apple"
	// fruitArray[1] = "orange"

	// Declare n Assign
	fruitArray := [2]string{"apple", "orange"}
	fruitSlice := []string{"apple", "orange", "grapes", "guava"}

	fmt.Println(fruitArray)
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
